package dbclients

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlClient struct {
	key string
	db  *sql.DB
	mu  *sync.Mutex
}

type mysqlTransaction struct {
	tx *sql.Tx
}

type mysqlExecResult struct {
	rowsAffected, lastInsertedID int64
}

func buildClient(config Config) (*mysqlClient, error) {
	c := &mysqlClient{
		key: config.Key,
		mu:  &sync.Mutex{},
	}

	if err := c.newConnection(config); err != nil {
		return nil, err
	}

	return c, nil
}

var sqlOpen = func(driver, conStr string) (*sql.DB, error) {
	db, err := sql.Open(driver, conStr)
	if err != nil {
		return nil, err
	}

	withConnectionRetry(func() error {
		err = db.Ping()
		return err
	})

	if err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}

func (c *mysqlClient) newConnection(config Config) error {
	if c.db != nil {
		return fmt.Errorf("dbclients: pool [%s] already has a connection", c.key)
	}

	conStr, err := c.buildConnectionString(config.Name, config.User, config.Pass, config.Host, config.Port, config.Query)
	if err != nil {
		return err
	}

	for i := 0; i < 60; i++ {
		c.db, err = sqlOpen("mysql", conStr)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		break
	}

	if err != nil {
		return err
	}

	// configs
	if config.MaxOpen > 0 {
		c.db.SetMaxOpenConns(config.MaxOpen)
	}

	if config.MaxOpenLife > 0 {
		c.db.SetConnMaxLifetime(config.MaxOpenLife)
	}

	if config.MaxIdle > 0 {
		c.db.SetMaxIdleConns(config.MaxIdle)
	}

	if config.MaxIdleLife > 0 {
		c.db.SetConnMaxIdleTime(config.MaxIdleLife)
	}

	log.Printf("dbclients: [%s.%s] connection open", config.Key, config.Name)
	return nil
}

func (c *mysqlClient) buildConnectionString(name, user, pass, host, port, query string) (string, error) {
	if user == "" {
		return "", errors.New("dbclients: empty user")
	}

	if pass == "" && user != "root" {
		return "", errors.New("dbclients: empty password")
	}

	if host == "" {
		return "", errors.New("dbclients: empty host")
	}

	return fmt.Sprintf(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, pass, host, port, name, query)), nil
}

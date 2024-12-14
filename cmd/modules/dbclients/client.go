package dbclients

import (
	"context"
	"fmt"
	"time"
)

type Client interface {
	GetKey() string
	Ping() error
	Query(ctx context.Context, query string, rowsMapper RowsMapper, args ...any) error
	Begin() (Transaction, error)
	Close() error
}

type Transaction interface {
	Exec(ctx context.Context, query string, args ...any) (Result, error)
	Commit() error
	Rollback() error
}

type Result interface {
	AffectedRows() int64
	InsertedID() int64
}

type RowsMapper func(func(dest ...any) error) error

type Config struct {
	Key         string
	Name        string
	Host        string
	Port        string
	User        string
	Pass        string
	Query       string
	MaxOpen     int
	MaxOpenLife time.Duration
	MaxIdle     int
	MaxIdleLife time.Duration
}

func Build(config Config) (Client, error) {
	if c, ok := findInCache(config.Key); ok {
		return c, nil
	}

	c, err := buildClient(config)
	if err != nil {
		return nil, err
	}

	return addToCache(config.Key, c), nil
}

func Get(key string) Client {
	if c, ok := findInCache(key); ok {
		return c
	}

	panic(fmt.Sprintf("dbclients: [%s] not found", key))
}

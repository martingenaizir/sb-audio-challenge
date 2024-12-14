package dbclients

import (
	"context"
	"database/sql"
	"log"
	"strings"
)

func (c *mysqlClient) GetKey() string {
	return c.key
}

func (c *mysqlClient) Ping() (err error) {
	withConnectionRetry(func() error {
		err = c.db.Ping()
		return err
	})

	return err
}

func (c *mysqlClient) Query(ctx context.Context, query string, rowsMapper RowsMapper, args ...any) (err error) {
	var rows *sql.Rows

	withConnectionRetry(func() error {
		rows, err = c.db.QueryContext(ctx, query, args...)
		return err
	})

	if err != nil {
		return err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		if err = rowsMapper(rows.Scan); err != nil {
			return err
		}
	}

	return rows.Err()
}

func (c *mysqlClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	removeFromCache(c.key)
	return c.db.Close()
}

func (c *mysqlClient) Begin() (Transaction, error) {
	tx, err := c.db.Begin()

	return &mysqlTransaction{
		tx: tx,
	}, err
}

func (m *mysqlTransaction) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	res, err := m.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	lastID, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	rowsAff, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &mysqlExecResult{
		rowsAffected:   rowsAff,
		lastInsertedID: lastID,
	}, nil
}

func (m *mysqlTransaction) Commit() error {
	return m.tx.Commit()
}

func (m *mysqlTransaction) Rollback() error {
	return m.tx.Rollback()
}

func (m *mysqlExecResult) AffectedRows() int64 {
	return m.rowsAffected
}

func (m *mysqlExecResult) InsertedID() int64 {
	return m.lastInsertedID
}

func withConnectionRetry(fn func() error) {
	const _invalidConnection = "invalid connection"
	if err := fn(); err != nil && strings.Contains(err.Error(), _invalidConnection) {
		log.Println("dbclients: invalid connection, retrying")
		_ = fn()
	}
}

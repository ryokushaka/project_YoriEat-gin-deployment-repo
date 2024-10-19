package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	// Importing the pq package for its side-effects (e.g., registering the PostgreSQL driver).
	_ "github.com/lib/pq"
)

// Client represents a PostgreSQL client.
type Client struct {
	DB *sqlx.DB
}

// QueryRow executes a query that is expected to return at most one row.
func (c *Client) QueryRow(query string, args ...interface{}) *sqlx.Row {
	return c.DB.QueryRowx(query, args...)
}

// Exec executes a query without returning any rows.
func (c *Client) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.DB.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (c *Client) Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	return c.DB.Queryx(query, args...)
}

// Begin starts a new transaction.
func (c *Client) Begin() (*sqlx.Tx, error) {
	return c.DB.Beginx()
}

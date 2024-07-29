package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Client struct {
	DB *sqlx.DB
}

func (c *Client) QueryRow(query string, args ...interface{}) *sqlx.Row {
	return c.DB.QueryRowx(query, args...)
}

func (c *Client) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.DB.Exec(query, args...)
}

func (c *Client) Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	return c.DB.Queryx(query, args...)
}

func (c *Client) Begin() (*sqlx.Tx, error) {
	return c.DB.Beginx()
}
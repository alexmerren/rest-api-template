package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"golang-api-template/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

type ConnProvider interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
}

type Datastore struct {
	db ConnProvider
}

// nolint:ineffassign,staticcheck // This allows us to check if any of them have an error, and return that error
// https://go.dev/doc/effective_go#redeclaration
func ProvideDatastore(context context.Context, config config.ConfigInterface) (*Datastore, error) {
	databaseUser, err := config.GetString("Database.Username")
	databasePass, err := config.GetString("Database.Password")
	databasePort, err := config.GetInt("Database.Port")
	databaseName, err := config.GetString("Database.Name")
	host, err := config.GetString("Host")
	if err != nil {
		return nil, err
	}

	configString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		databaseUser,
		databasePass,
		host,
		databasePort,
		databaseName,
	)
	db, err := sql.Open("mysql", configString)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(context)
	if err != nil {
		return nil, err
	}

	return &Datastore{
		db: db,
	}, nil
}

func (d *Datastore) CloseDB() error {
	return d.db.Close()
}

func (d *Datastore) QueryContext(context context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := d.db.QueryContext(context, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Datastore) QueryRowContext(context context.Context, query string, args ...interface{}) (*sql.Row, error) {
	row := d.db.QueryRowContext(context, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func (d *Datastore) ExecContext(context context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx, err := d.db.BeginTx(context, nil /* opts */)
	if err != nil {
		return nil, err
	}

	result, err := tx.ExecContext(context, query, args...)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return nil, rollbackErr
		}
	}
	return result, nil
}

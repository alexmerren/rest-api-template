package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"golang-api-template/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	InsertContact(context context.Context, contact *Contact) error
	GetContact(context context.Context, id string) (*Contact, error)
	GetAllContacts(context context.Context) ([]*Contact, error)
	UpdateContact(context context.Context, contact *Contact) error
	DeleteContact(context context.Context, id string) error
}

type ConnProvider interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
}

type Datastore struct {
	db ConnProvider
}

func ProvideDatastore(context context.Context, config config.Config) (*Datastore, error) {
	databaseUser, userErr := config.GetString("Database.Username")
	if userErr != nil {
		return nil, userErr
	}

	databasePass, passErr := config.GetString("Database.Password")
	if passErr != nil {
		return nil, passErr
	}

	databasePort, portErr := config.GetInt("Database.Port")
	if portErr != nil {
		return nil, portErr
	}

	databaseName, nameErr := config.GetString("Database.Name")
	if nameErr != nil {
		return nil, nameErr
	}

	host, hostErr := config.GetString("Host")
	if hostErr != nil {
		return nil, hostErr
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

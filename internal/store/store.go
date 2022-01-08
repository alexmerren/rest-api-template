package store

import (
	"context"
	"database/sql"
	"fmt"
	"golang-api-template/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db *sql.DB
}

// nolint:ineffassign,staticcheck // This allows us to check if any of them have an error, and return that error
// https://go.dev/doc/effective_go#redeclaration
func NewStore(context context.Context, config config.ConfigInterface) (*Store, error) {
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

	return &Store{
		db: db,
	}, nil
}

func (s *Store) CloseDB() error {
	return s.db.Close()
}

func (s *Store) QueryContext(context context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.db.QueryContext(context, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *Store) QueryRowContext(context context.Context, query string, args ...interface{}) *sql.Row {
	row := s.db.QueryRowContext(context, query, args...)
	return row
}

func (s *Store) ExecContext(context context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx, err := s.db.BeginTx(context, nil /* opts */)
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

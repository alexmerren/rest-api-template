package store

import (
	"database/sql"
	"fmt"
	"golang-api-template/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db *sql.DB
}

func NewStore() (*Store, error) {
	config := config.GetConfig()
	configString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.Database.Username,
		config.Database.Password,
		config.Host,
		config.Database.Port,
		config.Database.Name,
	)
	db, err := sql.Open("mysql", configString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *Store) QueryRow(query string, args ...interface{}) *sql.Row {
	row := s.db.QueryRow(query, args...)
	return row
}

func (s *Store) Exec(query string, args ...interface{}) (sql.Result, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	result, err := tx.Exec(query, args...)
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

package store

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestNewStore_HappyPath(t *testing.T) {}

func TestNewStore_ConfigErrors(t *testing.T) {}

func TestNewStore_OpenError(t *testing.T) {}

func TestNewStore_PingError(t *testing.T) {}

func TestQueryContext_HappyPath(t *testing.T) {}

func TestQueryRowContext_HappyPath(t *testing.T) {}

func TestExecContext_HappyPath(t *testing.T) {}

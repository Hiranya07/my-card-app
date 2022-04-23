package repositories

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"my-card-app/db"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

type Account struct {
}

func TestRepository_CreateAccount(t *testing.T) {
	dbIns, mock := NewMock()
	db.DbInstance = dbIns

	query := "INSERT INTO accounts \\(document_number\\) VALUES \\(?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1234).WillReturnResult(sqlmock.NewResult(0, 1))

	err := NewRepo().CreateAccount(context.Background(), 12345)
	assert.Error(t, err)
}

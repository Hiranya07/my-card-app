package repositories

import (
	"context"
	"database/sql"
	"my-card-app/db"
	"my-card-app/models"
	"my-card-app/response"
	"net/http"
)

type IRepository interface {
	CreateAccount(ctx context.Context, acc_id int) error
	GetAccount(ctx context.Context, accountId int) (models.AccountDetails, error)
}

type Repository struct{}

func NewRepo() IRepository {
	return Repository{}
}

func (repo Repository) CreateAccount(ctx context.Context, accountId int) error {

	_, err := db.DbInstance.Exec("INSERT INTO accounts (document_number) VALUES (?)", accountId)
	if err != nil {
		return err
	}
	return nil
}

func (repo Repository) GetAccount(ctx context.Context, accountId int) (models.AccountDetails, error) {

	var acc models.AccountDetails
	row := db.DbInstance.QueryRow("SELECT * FROM accounts WHERE acc_id = ?", accountId)
	err := row.Scan(&acc.AccId, &acc.DocumentNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			return acc, response.NewError(err, "record not found", err.Error(), http.StatusNotFound)
		}
		return acc, response.NewError(err, "error during record fetch", err.Error(), http.StatusInternalServerError)
	}
	return acc, nil
}

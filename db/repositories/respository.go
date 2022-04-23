package repositories

import (
	"context"
	"my-card-app/db"
)

type IRepository interface {
	CreateAccount(ctx context.Context, acc_id int) error
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

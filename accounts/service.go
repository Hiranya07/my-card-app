package accounts

import (
	"context"
	"my-card-app/db/repositories"
)

type IAccountService interface {
	CreateAccount(ctx context.Context, accountId int) error
}

type AccountService struct {
	repo func() repositories.IRepository
}

func NewAccountService() IAccountService {
	return AccountService{
		repo: repositories.NewRepo,
	}
}

func (accServer AccountService) CreateAccount(ctx context.Context, accountId int) error {

	return accServer.repo().CreateAccount(ctx, accountId)
}

// func (accServer AccountService) createRecord(ctx context.Context, accountId int) error {

// 	_, err := db.DbInstance.Exec("INSERT INTO accounts (document_number) VALUES (?)", accountId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

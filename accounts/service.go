package accounts

import (
	"context"
	"my-card-app/db/repositories"
	"my-card-app/models"
)

type IAccountService interface {
	CreateAccount(ctx context.Context, accountId int) error
	GetAccount(ctx context.Context, accountId int) (models.AccountDetails, error)
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

func (accServer AccountService) GetAccount(ctx context.Context, accountId int) (models.AccountDetails, error) {

	return accServer.repo().GetAccount(ctx, accountId)
}

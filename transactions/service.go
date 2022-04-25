package transactions

import (
	"context"
	"my-card-app/models"
	"my-card-app/transactions/db/repositories"
)

type ITransaction interface {
	Create(ctx context.Context, trnx models.Transactions) (models.TransactionResponse, error)
}

type TransactionService struct {
	repo func() repositories.IRepository
}

func NewTransactionService() ITransaction {
	return TransactionService{
		repo: repositories.NewRepo,
	}
}

func (trx TransactionService) Create(ctx context.Context, trnx models.Transactions) (models.TransactionResponse, error) {

	return trx.repo().Create(ctx, trnx)
}

package transactions

import (
	"context"
	"fmt"
	"my-card-app/models"
	"my-card-app/transactions/db/repositories"
)

type ITransaction interface {
	Create(ctx context.Context, trnx models.Transactions) error
}

type TransactionService struct {
	repo func() repositories.IRepository
}

func NewTransactionService() ITransaction {
	return TransactionService{
		repo: repositories.NewRepo,
	}
}

func (trx TransactionService) Create(ctx context.Context, trnx models.Transactions) error {

	fmt.Println("trx", trnx)
	return trx.repo().Create(ctx, trnx)
}

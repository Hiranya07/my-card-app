package repositories

import (
	"context"
	"my-card-app/db"
	"my-card-app/models"
	"time"
)

type IRepository interface {
	Create(ctx context.Context, trnx models.Transactions) error
}

type Repository struct{}

func NewRepo() IRepository {
	return Repository{}
}

func (repo Repository) Create(ctx context.Context, trnx models.Transactions) error {

	_, err := db.DbInstance.Exec("INSERT INTO transactions (Account_Id,OperationType_Id,Amount,EventDate) VALUES (?,?,?,?)", trnx.AccountId, trnx.OperationTypeId, trnx.Amount, time.Now())
	if err != nil {
		return err
	}
	return nil
}

package repositories

import (
	"context"
	"my-card-app/db"
	"my-card-app/models"
	"time"
)

type IRepository interface {
	Create(ctx context.Context, trnx models.Transactions) (models.TransactionResponse, error)
}

type Repository struct{}

func NewRepo() IRepository {
	return Repository{}
}

func (repo Repository) Create(ctx context.Context, trnx models.Transactions) (models.TransactionResponse, error) {

	var resp models.TransactionResponse
	result, err := db.DbInstance.Exec("INSERT INTO transactions (Account_Id,OperationType_Id,Amount,EventDate) VALUES (?,?,?,?)", trnx.AccountId, trnx.OperationTypeId, trnx.Amount, time.Now())
	if err != nil {
		return resp, err
	}

	trxId, _ := result.LastInsertId()
	resp.TransactionId = int(trxId)

	return resp, nil
}

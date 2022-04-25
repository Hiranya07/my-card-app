package transactions

import (
	"my-card-app/models"
	"my-card-app/response"
	"my-card-app/validations"
	"net/http"
)

type Transaction struct {
	trxServ func() ITransaction
}

const (
	errorCantDecodeInputData      = `error_can't_decode_input`
	errorWhileCreatingTransaction = `error_while_creating_transaction`
)

func NewTransaction() Transaction {
	return Transaction{
		trxServ: NewTransactionService,
	}
}

func (trx Transaction) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var trxns models.Transactions

	err := validations.ExtractPayload(r, &trxns)
	if err != nil {
		responseError := response.NewError(err, errorCantDecodeInputData, err.Error(), http.StatusBadRequest)
		response.Error(ctx, w, responseError, false)
		return
	}

	resp, err := trx.trxServ().Create(ctx, trxns)
	if err != nil {
		responseError := response.NewError(err, errorWhileCreatingTransaction, err.Error(), http.StatusInternalServerError)
		response.Error(ctx, w, responseError, false)
		return
	}

	response.Response(w, http.StatusCreated, resp)

}

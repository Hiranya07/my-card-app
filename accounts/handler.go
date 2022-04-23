package accounts

import (
	"my-card-app/models"
	"my-card-app/response"
	"my-card-app/validations"
	"net/http"
)

type Account struct {
	acc_id   int
	accServe func() IAccountService
}

func NewAccount() Account {
	return Account{
		accServe: NewAccountService,
	}
}

const (
	errorCantDecodeInputData  = `error_can't_decode_input`
	errorWhileCreatingAccount = `error_while_creating_account`
)

func (acc Account) CreateAccount(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var acc_in models.AccountInput
	err := validations.ExtractPayload(r, &acc_in)

	if err != nil {
		responseError := response.NewError(err, errorCantDecodeInputData, err.Error(), http.StatusBadRequest)
		response.Error(ctx, w, responseError, false)
		return
	}

	err = acc.accServe().CreateAccount(ctx, acc_in.AccId)
	if err != nil {
		responseError := response.NewError(err, errorWhileCreatingAccount, err.Error(), http.StatusInternalServerError)
		response.Error(ctx, w, responseError, false)
	}

	response.Response(w, http.StatusCreated, nil)

}

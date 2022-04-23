package accounts

import (
	"my-card-app/models"
	"my-card-app/response"
	"my-card-app/validations"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	errorCantDecodeInputData    = `error_can't_decode_input`
	errorWhileCreatingAccount   = `error_while_creating_account`
	errorWhileRetrievingAccount = `error_while_retreving_account`
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
		return
	}

	response.Response(w, http.StatusCreated, nil)

}

func (acc Account) GetAccount(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	accountIdStr := mux.Vars(r)["accountId"]

	accountId, _ := strconv.Atoi(accountIdStr)

	accDetails, err := acc.accServe().GetAccount(ctx, accountId)
	if err != nil {
		response.Error(ctx, w, err, false)
		return
	}

	response.Response(w, http.StatusOK, accDetails)

}

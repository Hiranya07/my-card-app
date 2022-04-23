package routes

import (
	"my-card-app/accounts"
	"my-card-app/transactions"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/accounts", accounts.NewAccount().CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{accountId}", accounts.NewAccount().GetAccount).Methods(http.MethodGet)

	router.HandleFunc("/transactions", transactions.NewTransaction().Create).Methods(http.MethodPost)
	return router
}

// func getAccounts(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("accounts", w, r)
// }

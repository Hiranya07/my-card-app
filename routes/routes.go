package routes

import (
	"my-card-app/accounts"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/accounts", accounts.NewAccount().CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{accountId}", accounts.NewAccount().GetAccount).Methods(http.MethodGet)
	return router
}

// func getAccounts(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("accounts", w, r)
// }

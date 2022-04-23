package main

import (
	"fmt"
	db "my-card-app/db"
	routes "my-card-app/routes"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {

	db.NewDbInstance()
	router := routes.NewRouter()
	middlewareManager := negroni.New()

	middlewareManager.UseHandler(router)

	err := http.ListenAndServe(":8080", middlewareManager)

	if err != nil {
		fmt.Println("error occurred", err)
	}
}

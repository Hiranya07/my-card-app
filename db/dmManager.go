package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DbInstance *sql.DB

func NewDbInstance() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "Hiranya",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "cardapp",
	}
	// Get a database handle.
	var err error
	DbInstance, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DbInstance.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

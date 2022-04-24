package db

import (
	"database/sql"
	"fmt"
	"log"

	conf "my-card-app/config"

	"github.com/go-sql-driver/mysql"
)

var DbInstance *sql.DB

func NewDbInstance() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: conf.Config.DbPassWord,
		Net:    "tcp",
		Addr:   conf.Config.DbAddr,
		DBName: conf.Config.DbName,
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

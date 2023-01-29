package model

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func ConnectDb() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "mysql",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "recordings",
	}
	// Get a database handle.
	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")
}

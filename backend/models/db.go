package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB(port string) bool {
	connStr := "host=localhost port=" + port + " user=postgres password=postgres dbname=bost sslmode=disable fallback_application_name=go_bost"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return false
	}

	DB = db

	return true
}

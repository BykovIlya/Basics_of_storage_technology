package models

import (
	"Basics_of_storage_technology/backend/utils"
	"database/sql"
	"fmt"
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

func Migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS movies(
        id BIGSERIAL PRIMARY KEY NOT NULL,
		title VARCHAR(255) NOT NULL,
		director VARCHAR(255) NOT NULL,
		year INTEGER,
		length INTEGER,
		studio VARCHAR(255) NOT NULL
    );
	CREATE TABLE IF NOT EXISTS directors(
        id BIGSERIAL PRIMARY KEY NOT NULL,
		name VARCHAR(255) NOT NULL,
		age INTEGER,
		gender VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL
    );
	CREATE TABLE IF NOT EXISTS boxoffice(
        id BIGSERIAL PRIMARY KEY NOT NULL,
		movie VARCHAR(255) NOT NULL,
		domestic_sales INTEGER,
		international_sales INTEGER
    );
	CREATE TABLE IF NOT EXISTS studios(
        id BIGSERIAL PRIMARY KEY NOT NULL,
		name VARCHAR(255) NOT NULL,
		year INTEGER,
		all_films INTEGER
    );
	`

	_, err := db.Exec(sql)

	if err != nil {
		utils.CheckErr(err)
		return
	}

	fmt.Println("Migrations successfull ended")
}

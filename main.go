package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "super"
	dbname   = "tours"
)

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO test.connection_test(id, text) VALUES($1, $2)", 1, "Test Connection")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")
}

func main() {
	initDB()
}

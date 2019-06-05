package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "atreya"
	dbname = "atreya"
)

func hellopostgres() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully Connected!")

	sqlStatement := `insert into users (age, email, first_name, last_name) values ($1, $2, $3, $4) returning id`
	id := 0
	err = db.QueryRow(sqlStatement, 30, "jo@bo.com", "jo", "bo").Scan(&id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("new id is: ", id)
}

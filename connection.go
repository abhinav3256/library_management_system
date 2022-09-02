package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://vgjvojuiaupsfk:3fae18ef1d97d7abc31f38e6a4135e3fe154fee347ac12f0633e2c139788f85b@ec2-52-73-155-171.compute-1.amazonaws.com:5432/ddaene760ljqmv"
)

func CreateDBConnection() {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	} else {
		fmt.Println("connected")
	}
	// defer DB.Close()
}

package main

import (
	"database/sql"
	//"encoding/csv"
)

var (
	DB *sql.DB
)

func main() {
	createDBConnection()

	importCSV()
}

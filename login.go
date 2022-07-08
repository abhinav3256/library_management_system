package main

import (
	"fmt"
	"log"
)

func login(reqBody User) {
	var count int
	row := DB.QueryRow("SELECT COUNT(*) FROM users")
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)

}

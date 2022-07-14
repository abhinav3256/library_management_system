package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func searchHandler(c *gin.Context) {
	name, ok := c.Params.Get("name")

	if ok == false {
		res := gin.H{
			"error": "Book name is missing",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	data := search(name)

	res := gin.H{
		"Books": data,
	}

	c.JSON(http.StatusBadRequest, res)
	return

}

func search(q string) []Book {
	Data := []Book{}

	SQL := "SELECT book_id ,book_name, book_author, book_cover_image,currrent_stock from books WHERE book_name LIKE $1"

	rows, err := DB.Query(SQL, "%"+q+"%")

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return Data
	}
	defer rows.Close()
	book := Book{}

	for rows.Next() {
		rows.Scan(&book.Book_id, &book.Book_name, &book.Book_author, &book.Book_cover_image, &book.Current_stock)

		//log.Fatal(err)
		Data = append(Data, book)

	}

	fmt.Println(Data)
	return Data

}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func orderHandler(c *gin.Context) {

	fmt.Println("2nd")
	email, _ := c.Cookie("email")
	fmt.Println(email)
	user_data := getUserByEmail(email)
	res := gin.H{
		"data": user_data,
	}
	c.JSON(http.StatusBadRequest, res)
	return
}

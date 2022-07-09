package main

import (
	"database/sql"
	//"encoding/csv"

	"github.com/gin-gonic/gin"
)

var (
	DB *sql.DB
)

type User struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email" binding:"email"`
	Password   string `json:"password" binding:"required,min=8,alphanum"`
}

func main() {
	createDBConnection()

	//importCSV()

	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func setupRoutes(r *gin.Engine) {

	r.POST("signup/admin", Signup)
	r.POST("signup/user", Signup)
	r.POST("login", login)
	r.POST("logout", logout)

}

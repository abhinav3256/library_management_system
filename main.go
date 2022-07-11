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
	Type       string `json:"type"`
}

type Order struct {
	ID                 int    `json:"id"`
	Book_id            string `json:"book_id"`
	Student_id         int    `json:"student_id"`
	Issue_date         string `json:"issue_date"`
	Return_date        string `json:"return_date"`
	Actual_return_date string `json:"actual_return_date"`
	Fine               string `json:"fine"`
	Approved           string `json:"approved"`
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

	r.GET("/orders", isLogin(), orderHandler)

	r.POST("/order", isLogin(), orderRequest)

}

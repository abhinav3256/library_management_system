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
	Type       int    `json:"type"`
}

type Orders struct {
	ID               int    `json:"id"`
	Book_id          string `json:"book_id"`
	Book_name        string `json:"book_name"`
	Book_author      string `json:"book_author"`
	Book_cover_image string `json:"book_cover_image"`
	Student_id       int    `json:"student_id"`
	Order_request_id int    `json:"order_request_id"`
	Issue_date       string `json:"issue_date"`
	Return_date      string `json:"return_date"`
	Fine             string `json:"fine"`
	Return_status    string `json:"return_status"`
}
type Order struct {
	ID                 int    `json:"id"`
	OrderRequestId     int    `json:"order_request_id" binding:"required"`
	Book_id            string `json:"book_id"`
	Student_id         int    `json:"student_id"`
	Issue_date         string `json:"issue_date"`
	Return_date        string `json:"return_date"`
	Actual_return_date string `json:"actual_return_date"`
	Fine               string `json:"fine"`
	Approved           string `json:"approved" binding:"required"`
}

type Order_request struct {
	ID                 int    `json:"id"`
	Book_id            string `json:"book_id"`
	Student_id         int    `json:"student_id"`
	Issue_date         string `json:"issue_date"`
	Return_date        string `json:"return_date"`
	Actual_return_date string `json:"actual_return_date"`
	Fine               string `json:"fine"`
	Approved           string `json:"approved"`
}

type Book struct {
	Book_id          string `json:"book_id"`
	Current_stock    int    `json:"current_stock"`
	Book_name        string `json:"book_name"`
	Book_author      string `json:"book_author"`
	Book_cover_image string `json:"book_cover_image"`
}

func main() {
	CreateDBConnection()

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

// 	r.GET("/orders", isLogin(), orderHandler)
// 	r.GET("/orders/history", isLogin(), orderHistoryHandler)

// 	r.POST("/order", isLogin(), orderRequest)

// 	r.POST("/order/return", isLogin(), order_return)

// 	r.GET("/search/:name", isLogin(), searchHandler)

// 	//ADMIN
// 	r.POST("/return/approve", isAdmin(), ReturnAccept)

// 	r.POST("/order/approve", isAdmin(), Order_approve)

// 	r.GET("/order/return", isAdmin(), OrderReturn)
	
	
	
	
	
	
	
	
	
	r.GET("/orders", orderHandler)
	r.GET("/orders/history", orderHistoryHandler)

	r.POST("/order",  orderRequest)

	r.POST("/order/return",  order_return)

	r.GET("/search/:name",  searchHandler)

	//ADMIN
	r.POST("/return/approve",  ReturnAccept)

	r.POST("/order/approve",  Order_approve)

	r.GET("/order/return",  OrderReturn)

}

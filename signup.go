package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func Signup(c *gin.Context) {
	reqBody := User{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": parseError(err),
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, result_err := InsertInDB(reqBody)
	if result_err != "" {
		res := gin.H{
			"error":   result_err,
			"success": result,
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := gin.H{
		"success": true,
		"result":  result,
		"message": "Signup Successfull",
	}
	c.JSON(http.StatusOK, res)
	return
}

func InsertInDB(reqbody User) (bool, string) {
	var result = true
	var err_responce = ""
	sqlStatement := `
INSERT INTO users(first_name, last_name,email, password, type)
VALUES ($1, $2, $3, $4,$5)`
	_, err2 := DB.Exec(sqlStatement, reqbody.First_name, reqbody.Last_name, reqbody.Email, reqbody.Password, 2)
	fmt.Println(err2)
	if err2 != nil {
		err := UniqueViolation(err2)
		if err != nil {
			fmt.Printf("%#v", err)
			return false, err.Detail
		}
		//log.Fatal("ERror in insert: ", err2)
		err_responce = "Something went wrong"
		result = false
		return result, err_responce
	}
	//fmt.Println(user)
	return result, err_responce
}
func UniqueViolation(err error) *pq.Error {
	if pqerr, ok := err.(*pq.Error); ok &&
		pqerr.Code == "23505" {
		return pqerr
	}
	return nil
}

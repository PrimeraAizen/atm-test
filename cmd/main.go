package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var operations = make(chan Operation)

type OperationRequest struct {
	Amount float64 `json:"amount"`
}

func createAccount(c *gin.Context) {
	account := NewAccount()
	c.JSON(http.StatusOK, gin.H{"account_id": account.ID})
}

func deposit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req OperationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(req)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	op := Operation{
		AccountID: id,
		Amount:    req.Amount,
		Type:      "deposit",
		Response:  make(chan error),
	}

	operations <- op
	err = <-op.Response
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func withdraw(c *gin.Context) {
	// withdraw money from account
	id, _ := strconv.Atoi(c.Param("id"))
	var req OperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	op := Operation{
		AccountID: id,
		Amount:    req.Amount,
		Type:      "withdraw",
		Response:  make(chan error),
	}

	operations <- op
	err := <-op.Response
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func getBalance(c *gin.Context) {
	// Get account balance
	id, _ := strconv.Atoi(c.Param("id"))

	op := Operation{
		AccountID: id,
		Type:      "balance",
		Response:  make(chan error),
	}

	operations <- op
	err := <-op.Response
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := Accounts[id]
	c.JSON(http.StatusOK, gin.H{"balance": account.GetBalance()})
}

func main() {
	go HandleOperations(operations)

	r := gin.Default()
	r.POST("/accounts", createAccount)
	r.POST("/accounts/:id/deposit", deposit)
	r.POST("/accounts/:id/withdraw", withdraw)
	r.GET("/accounts/:id/balance", getBalance)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}

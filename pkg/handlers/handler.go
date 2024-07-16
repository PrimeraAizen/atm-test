package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	account := router.Group("/account")
	{
		account.POST("/accounts", h.createAccount)
		account.POST("/{id}/deposit", h.deposit)
		account.POST("/{id}/withdraw", h.withdraw)
		account.GET("/{id}/balance", h.getBalance)
	}
	return router
}

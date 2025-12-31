package router

import (
	"net/http"

	"github.com/Rafiur/wallet_app/internal/handler"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, authenticate echo.MiddlewareFunc, mainHandler *handler.Handler) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "ok",
			"message": "Wallet App API is running",
			"version": "0.1.0",
		})
	})

	api := e.Group("/api/v1")
	api.Use(authenticate)

	// Accounts CRUD
	api.POST("/accounts", mainHandler.CreateAccount)
	api.GET("/accounts", mainHandler.ListAccounts)
	api.GET("/accounts/:id", mainHandler.GetAccount)
	api.PUT("/accounts/:id", mainHandler.UpdateAccount)
	api.DELETE("/accounts/:id", mainHandler.DeleteAccount)
	api.DELETE("/accounts", mainHandler.DeleteAccount) // supports bulk delete via JSON body

	// Users
	api.POST("/users", mainHandler.CreateUser)
	api.GET("/users", mainHandler.ListUsers)
	api.GET("/users/:id", mainHandler.GetUser)
	api.PUT("/users/:id", mainHandler.UpdateUser)
	api.DELETE("/users/:id", mainHandler.DeleteUser)
	api.DELETE("/users", mainHandler.DeleteUser)

	// Transactions
	api.POST("/transactions", mainHandler.CreateTransaction)
	api.GET("/transactions", mainHandler.ListTransactions)
	api.GET("/transactions/:id", mainHandler.GetTransaction)
	api.PUT("/transactions/:id", mainHandler.UpdateTransaction)
	api.DELETE("/transactions/:id", mainHandler.DeleteTransaction)
	api.DELETE("/transactions", mainHandler.DeleteTransaction)

	// Transfers
	api.POST("/transfers", mainHandler.CreateTransfer)
	api.GET("/transfers", mainHandler.ListTransfers)
	api.GET("/transfers/:id", mainHandler.GetTransfer)
	api.PUT("/transfers/:id", mainHandler.UpdateTransfer)
	api.DELETE("/transfers/:id", mainHandler.DeleteTransfer)
	api.DELETE("/transfers", mainHandler.DeleteTransfer)

	// Sessions
	api.POST("/sessions", mainHandler.CreateSession)
	api.GET("/sessions", mainHandler.ListSessions) // supports query by refresh_token
	api.GET("/sessions/:id", mainHandler.GetSession)
	api.DELETE("/sessions/:id", mainHandler.DeleteSession)
	api.DELETE("/sessions", mainHandler.DeleteSession)
}

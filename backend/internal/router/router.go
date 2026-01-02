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

	// Auth routes (no auth required)
	e.POST("/api/v1/login", mainHandler.Login)
	e.POST("/api/v1/refresh", mainHandler.RefreshToken)

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

	// Budgets
	api.POST("/budgets", mainHandler.CreateBudget)
	api.GET("/budgets", mainHandler.ListBudgets) // requires user_id
	api.GET("/budgets/:id", mainHandler.GetBudget)
	api.PUT("/budgets/:id", mainHandler.UpdateBudget)
	api.DELETE("/budgets/:id", mainHandler.DeleteBudget)

	// Currencies
	api.POST("/currencies", mainHandler.CreateCurrency)
	api.GET("/currencies", mainHandler.ListCurrencies)
	api.GET("/currencies/:code", mainHandler.GetCurrency)
	api.PUT("/currencies/:code", mainHandler.UpdateCurrency)
	api.DELETE("/currencies/:code", mainHandler.DeleteCurrency)

	// Expense Categories
	api.POST("/expense-categories", mainHandler.CreateExpenseCategory)
	api.GET("/expense-categories", mainHandler.ListExpenseCategories)
	api.GET("/expense-categories/:id", mainHandler.GetExpenseCategory)
	api.PUT("/expense-categories/:id", mainHandler.UpdateExpenseCategory)
	api.DELETE("/expense-categories/:id", mainHandler.DeleteExpenseCategory)
	api.DELETE("/expense-categories", mainHandler.DeleteExpenseCategory)

	// Banks
	api.POST("/banks", mainHandler.CreateBank)
	api.GET("/banks", mainHandler.ListBanks) // requires user_id
	api.GET("/banks/:id", mainHandler.GetBank)
	api.PUT("/banks/:id", mainHandler.UpdateBank)
	api.DELETE("/banks/:id", mainHandler.DeleteBank)

	// Investments
	api.POST("/investments", mainHandler.CreateInvestment)
	api.GET("/investments", mainHandler.ListInvestments) // requires user_id
	api.GET("/investments/:id", mainHandler.GetInvestment)
	api.PUT("/investments/:id", mainHandler.UpdateInvestment)
	api.DELETE("/investments/:id", mainHandler.DeleteInvestment)

	// Recurring Transactions
	api.POST("/recurring-transactions", mainHandler.CreateRecurringTransaction)
	api.GET("/recurring-transactions", mainHandler.ListRecurringTransactions) // requires user_id
	api.GET("/recurring-transactions/:id", mainHandler.GetRecurringTransaction)
	api.PUT("/recurring-transactions/:id", mainHandler.UpdateRecurringTransaction)
	api.DELETE("/recurring-transactions/:id", mainHandler.DeleteRecurringTransaction)

	// Account Currencies
	api.POST("/account-currencies", mainHandler.CreateAccountCurrency)
	api.GET("/account-currencies", mainHandler.ListAccountCurrencies) // requires account_id
	api.GET("/account-currencies/:id", mainHandler.GetAccountCurrency)
	api.PUT("/account-currencies/:id", mainHandler.UpdateAccountCurrency)
	api.DELETE("/account-currencies/:id", mainHandler.DeleteAccountCurrency)

	// Cash Flow Summary
	api.POST("/cash-flow-summaries", mainHandler.CreateCashFlowSummary)
	api.GET("/cash-flow-summaries", mainHandler.ListCashFlowSummaries) // requires user_id and period
	api.GET("/cash-flow-summaries/:id", mainHandler.GetCashFlowSummary)
	api.PUT("/cash-flow-summaries/:id", mainHandler.UpdateCashFlowSummary)
	api.DELETE("/cash-flow-summaries/:id", mainHandler.DeleteCashFlowSummary)
}

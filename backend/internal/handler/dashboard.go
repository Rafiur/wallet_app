package handler

import (
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/Rafiur/wallet_app/internal/domain/entity"
)

// GetDashboardSummary aggregates the caller's accounts and transactions for a
// given month (defaults to the current month) into the numbers the frontend
// dashboard needs: total balance, income/expense totals, a category
// breakdown for the expense/income charts, and the most recent transactions.
func (h *Handler) GetDashboardSummary(c echo.Context) error {
	ctx := c.Request().Context()
	userID := authUserID(c)

	now := time.Now()
	year, month := now.Year(), now.Month()
	if period := c.QueryParam("period"); period != "" {
		if t, err := time.Parse("2006-01", period); err == nil {
			year, month = t.Year(), t.Month()
		}
	}
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Second)

	accounts, err := h.accountService.List(ctx, &entity.FilterAccountListRequest{UserID: userID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	totalBalance := 0.0
	accountSummaries := make([]map[string]interface{}, 0, len(accounts))
	for _, a := range accounts {
		totalBalance += a.Balance
		accountSummaries = append(accountSummaries, map[string]interface{}{
			"id": a.ID, "name": a.Name, "type": a.Type, "currency": a.Currency, "balance": a.Balance,
		})
	}

	transactions, err := h.transactionService.List(ctx, &entity.FilterTransactionListRequest{
		UserID:    userID,
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	categories, err := h.expenseCategoryService.List(ctx, &entity.FilterExpenseCategoryListRequest{UserID: userID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	categoryNames := make(map[string]string, len(categories))
	for _, cat := range categories {
		categoryNames[cat.ID] = cat.Name
	}

	var incomeTotal, expenseTotal float64
	expenseByCategory := map[string]float64{}
	incomeByCategory := map[string]float64{}
	for _, t := range transactions {
		catName := "Uncategorized"
		if t.ExpenseCategoryID != nil && *t.ExpenseCategoryID != "" {
			if name, ok := categoryNames[*t.ExpenseCategoryID]; ok {
				catName = name
			} else {
				catName = "Unknown"
			}
		}
		amount := t.Amount
		if amount < 0 {
			amount = -amount
		}
		if t.TransactionType == "income" {
			incomeTotal += amount
			incomeByCategory[catName] += amount
		} else {
			expenseTotal += amount
			expenseByCategory[catName] += amount
		}
	}

	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].TransactionDate.After(transactions[j].TransactionDate)
	})
	recentLimit := 10
	if len(transactions) < recentLimit {
		recentLimit = len(transactions)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"period":              startDate.Format("2006-01"),
		"total_balance":       totalBalance,
		"accounts":            accountSummaries,
		"income_total":        incomeTotal,
		"expense_total":       expenseTotal,
		"net":                 incomeTotal - expenseTotal,
		"expense_by_category": categoryBreakdown(expenseByCategory),
		"income_by_category":  categoryBreakdown(incomeByCategory),
		"recent_transactions": transactions[:recentLimit],
	})
}

func categoryBreakdown(totals map[string]float64) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(totals))
	for name, amount := range totals {
		out = append(out, map[string]interface{}{"category": name, "amount": amount})
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i]["amount"].(float64) > out[j]["amount"].(float64)
	})
	return out
}

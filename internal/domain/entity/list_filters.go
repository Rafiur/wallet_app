package entity

import "time"

type FilterAccountListRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type FilterExpenseCategoryListRequest struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ParentCategoryID string `json:"parent_category_id"`
}

type FilterUserListRequest struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

type FilterTransactionListRequest struct {
	ID                string
	AccountID         string
	UserID            string
	ExpenseCategoryID string
	TransactionType   string
	Tags              []string

	StartDate time.Time
	EndDate   time.Time

	SearchBy  string // Search by TransactionName or Note
	SortBy    string // e.g., "transaction_date", "amount"
	SortOrder string // "asc", "desc"
}

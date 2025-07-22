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

	SearchBy string // Search by TransactionName or Note
	//SortBy    string // e.g., "transaction_date", "amount"
	//SortOrder string // "asc", "desc"
}

type FilterTransferListRequest struct {
	ID            string    // filter by transfer ID
	FromAccountID string    // filter by source account
	ToAccountID   string    // filter by destination account
	MinAmount     float64   // filter transfers with amount >= MinAmount
	MaxAmount     float64   // filter transfers with amount <= MaxAmount
	Currency      string    // filter by currency code
	Status        string    // filter by transfer status (e.g., "pending", "completed", "failed")
	StartDate     time.Time // filter transfers on or after this date
	EndDate       time.Time // filter transfers on or before this date
	NoteContains  string    // text search within the note field
	//SortBy        string    // e.g., "transfer_date", "amount"
	//SortOrder     string    // "asc" or "desc"
}

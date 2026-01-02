package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type Transaction struct {
	bun.BaseModel     `bun:"table:transactions"`
	ID                string     `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	AccountID         string     `bun:"account_id,notnull" json:"account_id"`
	TransactionName   string     `bun:"transaction_name,notnull" json:"transaction_name"`
	TransactionDate   time.Time  `bun:"transaction_date,default:current_timestamp" json:"transaction_date"`
	Amount            float64    `bun:"amount,notnull" json:"amount"` // Positive for income, negative for expense
	ExpenseCategoryID *string    `bun:"expense_category_id" json:"expense_category_id"`
	UserID            string     `bun:"user_id,notnull" json:"user_id"`
	Note              string     `bun:"note" json:"note"`
	Tags              []string   `bun:"tags,type:text[]" json:"tags"`
	TransactionType   string     `bun:"transaction_type,notnull" json:"transaction_type"` // 'income', 'expense'
	IsBillable        bool       `bun:"is_billable,default:false" json:"is_billable"`
	RecurringID       *string    `bun:"recurring_id" json:"recurring_id"` // Link if generated from recurring
	DeletedAt         *time.Time `bun:"deleted_at,soft_delete,nullzero" json:"deleted_at,omitempty"`
}

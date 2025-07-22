package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type Transaction struct {
	bun.BaseModel     `bun:"table:transaction"`
	ID                string    `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	AccountID         string    `bun:"account_id,notnull" json:"account_id"`
	TransactionName   string    `bun:"transaction_name,notnull" json:"transaction_name"`
	TransactionDate   time.Time `bun:"transaction_date,default:current_timestamp" json:"transaction_date"`
	Amount            float32   `bun:"amount" json:"amount"`
	ExpenseCategoryID string    `bun:"expense_category_id" json:"expense_category_id"`
	UserID            string    `bun:"user_id" json:"user_id"`
	Note              string    `bun:"note" json:"note"`
	Tags              []string  `bun:"tags,type:text[]" json:"tags"`
	TransactionType   string    `bun:"transaction_type" json:"transaction_type"`
}

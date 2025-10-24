package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type RecurringTransaction struct {
	bun.BaseModel     `bun:"table:recurring_transactions"`
	ID                string     `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	UserID            string     `bun:"user_id,notnull" json:"user_id"`
	AccountID         string     `bun:"account_id,notnull" json:"account_id"`
	Name              string     `bun:"name,notnull" json:"name"`                               // e.g., "Bi-monthly utility", "Annual bank fee"
	Amount            float64    `bun:"amount,notnull" json:"amount"`                           // Positive/negative as per type
	Frequency         string     `bun:"frequency,notnull" json:"frequency"`                     // e.g., 'monthly', 'bi-monthly', 'quarterly', 'semi-annual', 'annual'
	FrequencyInterval int        `bun:"frequency_interval,default:1" json:"frequency_interval"` // e.g., 2 for every 2 months
	StartDate         time.Time  `bun:"start_date,notnull" json:"start_date"`
	EndDate           *time.Time `bun:"end_date" json:"end_date"`                         // Optional end
	NextDueDate       time.Time  `bun:"next_due_date" json:"next_due_date"`               // For scheduling
	TransactionType   string     `bun:"transaction_type,notnull" json:"transaction_type"` // 'income', 'expense'
	ExpenseCategoryID *string    `bun:"expense_category_id" json:"expense_category_id"`
	Note              string     `bun:"note" json:"note"`
	IsActive          bool       `bun:"is_active,default:true" json:"is_active"`
	CreatedAt         time.Time  `bun:"created_at,default:current_timestamp"`
	UpdatedAt         time.Time  `bun:"updated_at,default:current_timestamp"`
	DeletedAt         *time.Time `bun:"deleted_at,soft_delete,nullzero"`
}

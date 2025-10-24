package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type Budget struct {
	bun.BaseModel     `bun:"table:budgets"`
	ID                string     `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	UserID            string     `bun:"user_id,notnull" json:"user_id"`
	ExpenseCategoryID *string    `bun:"expense_category_id" json:"expense_category_id"` // Per category or global (null)
	Period            string     `bun:"period,notnull" json:"period"`                   // 'monthly', 'quarterly', 'yearly'
	StartDate         time.Time  `bun:"start_date,notnull" json:"start_date"`
	EndDate           time.Time  `bun:"end_date" json:"end_date"`
	Amount            float64    `bun:"amount,notnull" json:"amount"` // Limit
	CreatedAt         time.Time  `bun:"created_at,default:current_timestamp"`
	UpdatedAt         time.Time  `bun:"updated_at,default:current_timestamp"`
	DeletedAt         *time.Time `bun:"deleted_at,soft_delete,nullzero"`
}

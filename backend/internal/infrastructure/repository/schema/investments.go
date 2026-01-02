package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type Investment struct {
	bun.BaseModel `bun:"table:investments"`
	ID            string     `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	UserID        string     `bun:"user_id,notnull" json:"user_id"`
	AccountID     *string    `bun:"account_id" json:"account_id"` // Optional link to bank account
	Type          string     `bun:"type,notnull" json:"type"`     // e.g., 'fdr', 'stocks', 'bonds'
	Amount        float64    `bun:"amount,notnull" json:"amount"`
	InterestRate  float64    `bun:"interest_rate" json:"interest_rate"`
	StartDate     time.Time  `bun:"start_date,notnull" json:"start_date"`
	MaturityDate  time.Time  `bun:"maturity_date" json:"maturity_date"`
	Status        string     `bun:"status,default:'active'" json:"status"` // e.g., 'active', 'matured'
	Note          string     `bun:"note" json:"note"`
	CreatedAt     time.Time  `bun:"created_at,default:current_timestamp"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero"`
}

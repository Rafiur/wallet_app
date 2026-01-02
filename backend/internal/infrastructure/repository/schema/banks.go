package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type Bank struct {
	bun.BaseModel `bun:"table:banks"`
	ID            string     `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	UserID        string     `bun:"user_id,notnull" json:"user_id"`
	Name          string     `bun:"name,notnull" json:"name"` // e.g., "Bank of America"
	Branch        string     `bun:"branch" json:"branch"`
	AccountNumber string     `bun:"account_number,unique" json:"account_number"`
	AccountID     string     `bun:"account_id,notnull" json:"account_id"`         // Links to Account
	AnnualCharge  float64    `bun:"annual_charge,default:0" json:"annual_charge"` // For annual fees; can generate recurring expense
	CreatedAt     time.Time  `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:"updated_at,default:current_timestamp"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero"`
}

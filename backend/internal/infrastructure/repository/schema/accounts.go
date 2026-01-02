package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts"`
	ID            string     `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	UserID        string     `bun:"user_id,notnull" json:"user_id"`
	Name          string     `bun:"name,notnull" json:"name"`
	Type          string     `bun:"type,notnull" json:"type"` // e.g., 'bank', 'wallet', 'cash', 'investment'
	Balance       float64    `bun:"balance,default:0" json:"balance"`
	Currency      string     `bun:"currency,notnull" json:"currency"`
	CreatedAt     time.Time  `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:"updated_at,default:current_timestamp"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero"`
}

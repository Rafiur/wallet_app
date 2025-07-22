package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts"`
	ID            string  `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	UserID        string  `bun:"user_id,notnull" json:"user_id"`
	Name          string  `bun:"name,notnull" json:"name"`
	Balance       float32 `bun:"balance" json:"balance"`
	Currency      string  `bun:"currency" json:"currency"`

	CreatedAt time.Time  `bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time  `bun:"updated_at,default:current_timestamp"`
	DeletedAt *time.Time `bun:"deleted_at,default:current_timestamp"`
}

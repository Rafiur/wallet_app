package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type Session struct {
	bun.BaseModel `bun:"table:sessions" json:"-"`
	ID            string    `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	UserID        string    `bun:"user_id,notnull,type:uuid" json:"user_id"`
	RefreshToken  string    `bun:"refresh_token,notnull,unique" json:"-"` // hashed or raw depending on setup; never serialized
	UserAgent     string    `bun:"user_agent" json:"user_agent"`
	IPAddress     string    `bun:"ip_address" json:"ip_address"`
	ExpiresAt     time.Time `bun:"expires_at,notnull" json:"expires_at"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
}

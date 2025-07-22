package schema

import "time"

type Session struct {
	ID           string    `bun:"id,pk,default:uuid_generate_v4(),type:uuid"`
	UserID       string    `bun:"user_id,notnull,type:uuid"`
	RefreshToken string    `bun:"refresh_token,notnull,unique"` // hashed or raw depending on setup
	UserAgent    string    `bun:"user_agent"`
	IPAddress    string    `bun:"ip_address"`
	ExpiresAt    time.Time `bun:"expires_at,notnull"`
	CreatedAt    time.Time `bun:"created_at,default:current_timestamp"`
}

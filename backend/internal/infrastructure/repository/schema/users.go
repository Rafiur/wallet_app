package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            string     `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	FullName      string     `bun:"full_name,notnull" json:"full_name"`
	Email         string     `bun:"email,notnull,unique" json:"email"`
	Password      string     `bun:"password,notnull" json:"-"`
	CreatedAt     time.Time  `bun:"created_at,default:current_timestamp"`
	DeletedAt     *time.Time `bun:"deleted_at"`
}

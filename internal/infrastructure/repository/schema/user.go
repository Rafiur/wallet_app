package schema

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            string `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	FullName      string `bun:"full_name,notnull" json:"full_name"`
	Email         string `bun:"email,notnull" json:"email"`
	Password      string `bun:"password,notnull" json:"password"`
}

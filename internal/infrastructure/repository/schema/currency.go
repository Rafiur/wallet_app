package schema

import "github.com/uptrace/bun"

type Currency struct {
	bun.BaseModel `bun:"table:currencies"`

	Code   string `bun:",pk"`      // e.g. "USD"
	Name   string `bun:",notnull"` // e.g. "US Dollar"
	Symbol string `bun:",notnull"` // e.g. "$"
}

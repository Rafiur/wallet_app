package schema

import "github.com/uptrace/bun"

type Currency struct {
	bun.BaseModel `bun:"table:currencies"`

	Code   string `bun:",pk" json:"code"`        // e.g. "USD"
	Name   string `bun:",notnull" json:"name"`   // e.g. "US Dollar"
	Symbol string `bun:",notnull" json:"symbol"` // e.g. "$"
}

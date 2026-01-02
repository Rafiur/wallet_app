package schema

import (
	"time"

	"github.com/uptrace/bun"
)

type Transfer struct {
	bun.BaseModel `bun:"table:transfers"`

	ID            string     `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	FromAccountID string     `bun:"from_account_id,notnull,type:uuid" json:"from_account_id"`
	ToAccountID   string     `bun:"to_account_id,notnull,type:uuid" json:"to_account_id"`
	Amount        float64    `bun:"amount,notnull" json:"amount"`
	Currency      string     `bun:"currency,notnull" json:"currency"`   // Source currency
	ExchangeRate  *float64   `bun:"exchange_rate" json:"exchange_rate"` // If cross-currency
	TransferDate  time.Time  `bun:"transfer_date,notnull,default:current_timestamp" json:"transfer_date"`
	Note          string     `bun:"note" json:"note"`
	Status        string     `bun:"status,default:'pending'" json:"status"` // e.g., 'pending', 'completed', 'failed'
	FromAccount   *Account   `bun:"rel:belongs-to,join:from_account_id=id" json:"from_account,omitempty"`
	ToAccount     *Account   `bun:"rel:belongs-to,join:to_account_id=id" json:"to_account,omitempty"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero" json:"deleted_at,omitempty"`
}

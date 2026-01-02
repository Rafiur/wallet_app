package schema

import "github.com/uptrace/bun"

type AccountCurrencies struct {
	bun.BaseModel              `bun:"table:account_currencies"`
	ID                         string   `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	AccountID                  string   `bun:"account_id,notnull" json:"account_id"`
	CurrencyCode               string   `bun:"currency_code" json:"currency_code"`
	BalanceAccordingToCurrency float32  `bun:"balance_according_to_currency" json:"balance_according_to_currency"`
	Account                    *Account `bun:"rel:belongs-to,join:account_id=id" json:"account,omitempty"`
}

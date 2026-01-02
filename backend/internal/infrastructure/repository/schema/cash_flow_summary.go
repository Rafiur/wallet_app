package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type CashFlowSummary struct {
	bun.BaseModel   `bun:"table:cash_flow_summaries"`
	ID              string    `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	UserID          string    `bun:"user_id,notnull" json:"user_id"`
	AccountID       *string   `bun:"account_id" json:"account_id"` // Per account or global
	Period          string    `bun:"period,notnull" json:"period"` // e.g., 'daily', 'monthly', 'yearly'
	StartDate       time.Time `bun:"start_date,notnull" json:"start_date"`
	Inflow          float64   `bun:"inflow,default:0" json:"inflow"`     // Sum of income
	Outflow         float64   `bun:"outflow,default:0" json:"outflow"`   // Sum of expenses
	NetFlow         float64   `bun:"net_flow,default:0" json:"net_flow"` // Inflow - Outflow
	BillableOutflow float64   `bun:"billable_outflow,default:0" json:"billable_outflow"`
	CreatedAt       time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt       time.Time `bun:"updated_at,default:current_timestamp"`
}

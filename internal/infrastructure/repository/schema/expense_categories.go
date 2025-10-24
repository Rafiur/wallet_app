package schema

import (
	"github.com/uptrace/bun"
	"time"
)

type ExpenseCategory struct {
	bun.BaseModel    `bun:"table:expense_categories"`
	ID               string             `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	Name             string             `bun:"name,notnull" json:"name"`
	Type             string             `bun:"type,default:'expense'" json:"type"` // e.g., 'expense', 'income'
	UserID           *string            `bun:"user_id" json:"user_id"`             // Allow global vs. user-specific
	ParentCategoryID *string            `bun:"parent_category_id" json:"parent_category_id"`
	ParentCategory   *ExpenseCategory   `bun:"rel:belongs-to,join:parent_category_id=id" json:"parent_category"`
	SubCategories    []*ExpenseCategory `bun:"rel:has-many,join:id=parent_category_id" json:"sub_categories"`
	CreatedAt        time.Time          `bun:"created_at,default:current_timestamp"`
	DeletedAt        *time.Time         `bun:"deleted_at"`
}

package schema

import "github.com/uptrace/bun"

type ExpenseCategory struct {
	bun.BaseModel    `bun:"table:user_expense_category"`
	ID               string           `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	Name             string           `bun:"name,notnull" json:"name"`
	ParentCategoryID string           `bun:"parent_category_id" json:"parent_category_id"`
	ParentCategory   *ExpenseCategory `bun:"rel:belongs-to,join:id=parent_category_id" json:"parent_category"`
}

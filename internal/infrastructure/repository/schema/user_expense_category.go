package schema

import "github.com/uptrace/bun"

type UserExpenseCategory struct {
	bun.BaseModel    `bun:"table:user_expense_category"`
	ID               string               `bun:"id,pk,default:uuid_generate_v4(),type:uuid" json:"id"`
	Name             string               `bun:"name,notnull" json:"name"`
	UserID           string               `bun:"user_id,notnull" json:"user_id"`
	ParentCategoryID string               `bun:"parent_category_id" json:"parent_category_id"`
	ParentCategory   *UserExpenseCategory `bun:"rel:belongs-to,join:id=parent_category_id"`
}

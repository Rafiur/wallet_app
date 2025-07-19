package entity

type FilterAccountListRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type FilterExpenseCategoryListRequest struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ParentCategoryID string `json:"parent_category_id"`
}

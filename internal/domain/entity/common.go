package entity

type CommonDeleteReq struct {
	ID  string   `json:"id"`
	IDs []string `json:"ids"`
}

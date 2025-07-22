package entity

type CommonDeleteReq struct {
	ID  string   `json:"id,omitempty"`
	IDs []string `json:"ids,omitempty"`
}

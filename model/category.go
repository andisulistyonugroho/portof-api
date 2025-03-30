package model

type Category struct {
	ID        int    `json:"id" sql:"id"`
	PageId    int    `json:"page_id" sql:"page_id"`
	Title     string `json:"title" sql:"title"`
	ParentId  int    `json:"parent_id" sql:"parent_id"`
	IsActive  bool   `json:"is_active" sql:"is_active"`
	CreatedAt string `json:"created_at" sql:"created_at"`
	EditedAt  string `json:"edited_at" sql:"edited_At"`
}

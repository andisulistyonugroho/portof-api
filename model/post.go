package model

type Post struct {
	ID                int    `json:"id" sql:"id"`
	PageId            int    `json:"page_id" sql:"page_id"`
	Title             string `json:"title" sql:"title"`
	ShortDesc         string `json:"short_desc" sql:"short_desc"`
	TheBody           string `json:"the_body" sql:"the_body"`
	DisplayPictureUrl string `json:"display_picture_url" sql:"display_picture_url"`
	IsActive          bool   `json:"is_active" sql:"is_active"`
	CreatedAt         string `json:"created_at" sql:"created_at"`
	EditedAt          string `json:"edited_at" sql:"edited_At"`
}

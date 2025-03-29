package model

type User struct {
	ID int `json:"id" sql:"id"`
	Username string `json:"username" sql:"username"`
	Passw string `json:"passw" sql:"passw"`
	IsVerified string `json:"is_verified" sql:"is_verified"`
	CreatedAt string `json:"created_at" sql:"created_at"`
	EditedAt string `json:"edited_at" sql:"edited_at"`
}
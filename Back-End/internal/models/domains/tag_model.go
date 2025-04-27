package models

type Tag struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
	UserID  int    `json:"user_id"`
}

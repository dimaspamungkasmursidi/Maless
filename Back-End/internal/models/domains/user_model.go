package models

type User struct {
	ID        int    `json:"id"`
	GoogleId  string `json:"google_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
}

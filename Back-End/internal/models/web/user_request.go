package web

type UserRequest struct {
	GoogleId  string `json:"google_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

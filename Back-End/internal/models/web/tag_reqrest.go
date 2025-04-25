package web

type TagRequest struct {
	TagName string `json:"tag_name"`
	UserID  int    `json:"user_id"`
}

type TagResponse struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
	UserID  int    `json:"user_id"`
}

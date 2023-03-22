package handler

type UrlCreationRequest struct {
	InitialUrl string `json:"initial_url" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
}

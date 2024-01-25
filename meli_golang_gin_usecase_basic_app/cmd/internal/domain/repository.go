package domain

type Repository struct {
	URL string `json:"url" binding:"required"`
}

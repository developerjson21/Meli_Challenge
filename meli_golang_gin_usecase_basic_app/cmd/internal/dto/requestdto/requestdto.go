package requestdto

type RequestBody struct {
	URL string `json:"url" binding:"required"`
}

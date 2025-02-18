package requests

type URLCreateRequest struct {
	OriginalURL string `json:"originalURL" form:"originalURL"`
}

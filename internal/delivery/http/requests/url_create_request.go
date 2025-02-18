package requests

// URLCreateRequest represents request for creating short URL
type URLCreateRequest struct {
	OriginalURL string `json:"originalURL" form:"originalURL" example:"https://very-long-url.com/some/path"`
}

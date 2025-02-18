package responses

import "github.com/idmaksim/url-shortener-api/internal/domain/models"

type URLResponse struct {
	ShortURL    string `json:"shortURL"`
	OriginalURL string `json:"originalURL"`
}

func NewURLResponse(url *models.URL) *URLResponse {
	return &URLResponse{
		ShortURL:    url.ShortURL,
		OriginalURL: url.OriginalURL,
	}
}

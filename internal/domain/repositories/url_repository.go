package repositories

import (
	"github.com/idmaksim/url-shortener-api/internal/domain/models"
)

type URLRepository interface {
	Create(url *models.URL) (*models.URL, error)
	FindOneByShortURL(shortURL string) (*models.URL, error)
}

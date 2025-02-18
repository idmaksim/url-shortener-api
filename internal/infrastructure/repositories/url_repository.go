package repositories

import (
	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/idmaksim/url-shortener-api/internal/constants"
	"github.com/idmaksim/url-shortener-api/internal/db"
	"github.com/idmaksim/url-shortener-api/internal/domain/errors"
	"github.com/idmaksim/url-shortener-api/internal/domain/models"
)

type URLRepository struct {
	db *db.DB
}

func NewURLRepository(cfg *config.Config) *URLRepository {
	return &URLRepository{db: db.NewDB(cfg.Database)}
}

func (r *URLRepository) Create(url *models.URL) (*models.URL, error) {
	err := r.db.DB.Create(url).Error
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (r *URLRepository) FindOneByShortURL(shortURL string) (*models.URL, error) {
	var url models.URL
	err := r.db.DB.Where("short_url = ?", shortURL).First(&url).Error
	if err != nil {
		return nil, errors.ErrNotFound
	}

	url.AccessCount++

	if url.AccessCount >= constants.AccessCount {
		if err := r.db.DB.Delete(&url).Error; err != nil {
			return nil, err
		}
		return nil, errors.ErrNotFound
	}

	err = r.db.DB.Save(&url).Error
	if err != nil {
		return nil, err
	}
	return &url, nil
}

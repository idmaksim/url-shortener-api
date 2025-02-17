package repositories

import (
	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/idmaksim/url-shortener-api/internal/db"
)

type URLRepository struct {
	db *db.DB
}

func NewURLRepository(cfg *config.Config) *URLRepository {
	return &URLRepository{db: db.NewDB(cfg.Database)}
}

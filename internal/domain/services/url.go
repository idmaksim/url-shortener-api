package services

import (
	"context"

	"github.com/idmaksim/url-shortener-api/internal/domain/models"
)

type URLService interface {
	Create(ctx context.Context, url *models.URL) (*models.URL, error)
	Get(ctx context.Context, shortURL string) (*models.URL, error)
}

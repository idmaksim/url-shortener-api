package services

import (
	"context"

	"github.com/idmaksim/url-shortener-api/internal/domain/models"
)

type URLService struct {
}

func NewURLService() *URLService {
	return &URLService{}
}

func (s *URLService) Create(ctx context.Context, url *models.URL) (*models.URL, error) {
	return url, nil
}

func (s *URLService) Get(ctx context.Context, shortURL string) (*models.URL, error) {
	return nil, nil
}

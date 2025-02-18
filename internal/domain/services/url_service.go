package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/idmaksim/url-shortener-api/internal/delivery/http/requests"
	"github.com/idmaksim/url-shortener-api/internal/domain/models"
	domainRepositories "github.com/idmaksim/url-shortener-api/internal/domain/repositories"
	"github.com/idmaksim/url-shortener-api/internal/infrastructure/repositories"
)

type URLService struct {
	repo domainRepositories.URLRepository
	cfg  *config.Config
}

func NewURLService(cfg *config.Config) *URLService {
	return &URLService{
		repo: repositories.NewURLRepository(cfg),
		cfg:  cfg,
	}
}

func (s *URLService) Create(request requests.URLCreateRequest) (*models.URL, error) {
	url := &models.URL{
		OriginalURL: request.OriginalURL,
		ShortURL:    s.GenerateShortURL(request.OriginalURL),
	}

	newUrl, err := s.repo.Create(url)
	if err != nil {
		return nil, err
	}

	newUrl.ShortURL = fmt.Sprintf("%s/%s", s.cfg.Http.Host, newUrl.ShortURL)

	return newUrl, nil
}

func (s *URLService) Get(shortURL string) (*models.URL, error) {
	return s.repo.FindOneByShortURL(shortURL)

}

func (s *URLService) GenerateShortURL(originalURL string) string {
	return uuid.New().String()[:6]
}

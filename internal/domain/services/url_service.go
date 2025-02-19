package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/idmaksim/url-shortener-api/internal/constants"
	"github.com/idmaksim/url-shortener-api/internal/delivery/http/requests"
	"github.com/idmaksim/url-shortener-api/internal/domain/models"
	domainRepositories "github.com/idmaksim/url-shortener-api/internal/domain/repositories"
	"github.com/idmaksim/url-shortener-api/internal/infrastructure/cache"
	"github.com/idmaksim/url-shortener-api/internal/infrastructure/cache/redis"
	"github.com/idmaksim/url-shortener-api/internal/infrastructure/repositories"
	libRedis "github.com/redis/go-redis/v9"
)

type URLService struct {
	repo  domainRepositories.URLRepository
	cfg   *config.Config
	cache cache.Cache
}

func NewURLService(cfg *config.Config) *URLService {
	return &URLService{
		repo:  repositories.NewURLRepository(cfg),
		cfg:   cfg,
		cache: redis.NewRedisCache(cfg),
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
	cacheKey := fmt.Sprintf("url:%s", shortURL)

	val, err := s.cache.Get(cacheKey)

	if val != "" {
		return &models.URL{
			OriginalURL: val,
		}, nil
	}

	if err == libRedis.Nil {
		url, err := s.repo.FindOneByShortURL(shortURL)
		if err != nil {
			return nil, err
		}

		s.cache.Set(cacheKey, url.OriginalURL, constants.DefaultCacheTTL)

		return url, nil
	}

	return nil, err
}

func (s *URLService) GenerateShortURL(originalURL string) string {
	return uuid.New().String()[:6]
}

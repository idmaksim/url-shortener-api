package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/idmaksim/url-shortener-api/internal/constants"
	"github.com/idmaksim/url-shortener-api/internal/delivery/http/requests"
	"github.com/idmaksim/url-shortener-api/internal/domain/models"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockURLRepository struct {
	mock.Mock
}

func (m *MockURLRepository) Create(url *models.URL) (*models.URL, error) {
	args := m.Called(url)
	return args.Get(0).(*models.URL), args.Error(1)
}

func (m *MockURLRepository) FindOneByShortURL(shortURL string) (*models.URL, error) {
	args := m.Called(shortURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.URL), args.Error(1)
}

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Set(key string, value string, ttl time.Duration) error {
	args := m.Called(key, value, ttl)
	return args.Error(0)
}

func (m *MockCache) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func TestURLService_Create(t *testing.T) {
	mockRepo := new(MockURLRepository)
	mockCache := new(MockCache)
	cfg := &config.Config{
		Http: &config.HTTPConfig{
			Host: "http://localhost:8080",
		},
	}

	service := &URLService{
		repo:  mockRepo,
		cfg:   cfg,
		cache: mockCache,
	}

	tests := []struct {
		name        string
		request     requests.URLCreateRequest
		mockReturn  *models.URL
		mockError   error
		expectError bool
	}{
		{
			name: "Successfully created",
			request: requests.URLCreateRequest{
				OriginalURL: "https://example.com",
			},
			mockReturn: &models.URL{
				ShortURL:    "abc123",
				OriginalURL: "https://example.com",
			},
			mockError:   nil,
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.On("Create", mock.AnythingOfType("*models.URL")).Return(tc.mockReturn, tc.mockError)

			result, err := service.Create(tc.request)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Contains(t, result.ShortURL, cfg.Http.Host)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestURLService_Get(t *testing.T) {
	mockRepo := new(MockURLRepository)
	mockCache := new(MockCache)
	cfg := &config.Config{
		Http: &config.HTTPConfig{
			Host: "http://localhost:8080",
		},
	}

	service := &URLService{
		repo:  mockRepo,
		cfg:   cfg,
		cache: mockCache,
	}

	tests := []struct {
		name        string
		shortURL    string
		mockReturn  *models.URL
		mockError   error
		expectError bool
	}{
		{
			name:     "Successfully retrieved",
			shortURL: "abc123",
			mockReturn: &models.URL{
				ShortURL:    "abc123",
				OriginalURL: "https://example.com",
			},
			mockError:   nil,
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cacheKey := fmt.Sprintf("url:%s", tc.shortURL)
			mockCache.On("Get", cacheKey).Return("", redis.Nil)
			mockRepo.On("FindOneByShortURL", tc.shortURL).Return(tc.mockReturn, tc.mockError)
			mockCache.On("Set", cacheKey, tc.mockReturn.OriginalURL, constants.DefaultCacheTTL).Return(nil)

			result, err := service.Get(tc.shortURL)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.mockReturn.OriginalURL, result.OriginalURL)
			}

			mockRepo.AssertExpectations(t)
			mockCache.AssertExpectations(t)
		})
	}
}

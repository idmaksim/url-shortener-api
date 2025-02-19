package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/idmaksim/url-shortener-api/internal/delivery/http/requests"
	"github.com/idmaksim/url-shortener-api/internal/domain/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockURLService struct {
	mock.Mock
}

func (m *MockURLService) Create(request requests.URLCreateRequest) (*models.URL, error) {
	args := m.Called(request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.URL), args.Error(1)
}

func (m *MockURLService) Get(shortURL string) (*models.URL, error) {
	args := m.Called(shortURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.URL), args.Error(1)
}

func Test_URLHandler_Create(t *testing.T) {
	e := echo.New()
	mockService := new(MockURLService)
	handler := &URLHandler{urlService: mockService}

	tests := []struct {
		name           string
		request        requests.URLCreateRequest
		mockReturn     *models.URL
		mockError      error
		expectedStatus int
		expectedBody   string
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
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"shortURL":"abc123","originalURL":"https://example.com"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService.On("Create", tc.request).Return(tc.mockReturn, tc.mockError)

			requestBody, _ := json.Marshal(tc.request)
			req := httptest.NewRequest(http.MethodPost, "/url", bytes.NewBuffer(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.Create(c)

			if err != nil {
				assert.Equal(t, tc.mockError, err)
			} else {
				assert.Equal(t, tc.expectedStatus, rec.Code)
				assert.JSONEq(t, tc.expectedBody, rec.Body.String())
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestURLHandler_Get(t *testing.T) {
	e := echo.New()
	mockService := new(MockURLService)
	handler := &URLHandler{urlService: mockService}

	tests := []struct {
		name           string
		shortURL       string
		mockReturn     *models.URL
		mockError      error
		expectedStatus int
		expectedURL    string
	}{
		{
			name:     "Successfully retrieved",
			shortURL: "abc123",
			mockReturn: &models.URL{
				ShortURL:    "abc123",
				OriginalURL: "https://example.com",
			},
			mockError:      nil,
			expectedStatus: http.StatusTemporaryRedirect,
			expectedURL:    "https://example.com",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService.On("Get", tc.shortURL).Return(tc.mockReturn, tc.mockError)

			req := httptest.NewRequest(http.MethodGet, "/"+tc.shortURL, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("shortURL")
			c.SetParamValues(tc.shortURL)

			err := handler.Get(c)

			if err != nil {
				assert.Equal(t, tc.mockError, err)
			} else {
				assert.Equal(t, tc.expectedStatus, rec.Code)
				assert.Equal(t, tc.expectedURL, rec.Header().Get("Location"))
			}

			mockService.AssertExpectations(t)
		})
	}
}

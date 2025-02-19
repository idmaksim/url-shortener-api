package handlers

import (
	"net/http"

	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/idmaksim/url-shortener-api/internal/delivery/http/requests"
	"github.com/idmaksim/url-shortener-api/internal/delivery/http/responses"
	"github.com/idmaksim/url-shortener-api/internal/domain/services"
	"github.com/labstack/echo/v4"
)

type URLHandler struct {
	urlService *services.URLService
}

func NewURLHandler(cfg *config.Config) *URLHandler {
	return &URLHandler{
		urlService: services.NewURLService(cfg),
	}
}

// Create godoc
// @Summary Create short URL
// @Description Creates a short URL from the original URL
// @Tags urls
// @Accept json
// @Produce json
// @Param request body requests.URLCreateRequest true "URL to shorten"
// @Success 200 {object} responses.URLResponse
// @Failure 400 {object} errors.HttpError
// @Failure 500 {object} errors.HttpError
// @Router /url [post]
func (h *URLHandler) Create(c echo.Context) error {
	var request requests.URLCreateRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUrl, err := h.urlService.Create(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := responses.NewURLResponse(newUrl)

	return c.JSON(http.StatusOK, response)
}

func (h *URLHandler) Get(c echo.Context) error {
	shortURL := c.Param("shortURL")

	url, err := h.urlService.Get(shortURL)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, url.OriginalURL)
}

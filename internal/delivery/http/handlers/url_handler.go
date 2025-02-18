package handlers

import (
	"net/http"

	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/idmaksim/url-shortener-api/internal/delivery/http/requests"
	"github.com/idmaksim/url-shortener-api/internal/domain/errors"
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

func (h *URLHandler) Create(c echo.Context) error {
	var request requests.URLCreateRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUrl, err := h.urlService.Create(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newUrl)
}

func (h *URLHandler) Get(c echo.Context) error {
	shortURL := c.Param("shortURL")

	url, err := h.urlService.Get(shortURL)
	if err != nil {
		if err == errors.ErrNotFound {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, url)
}

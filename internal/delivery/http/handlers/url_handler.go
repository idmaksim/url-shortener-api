package handlers

import (
	"net/http"

	domainServices "github.com/idmaksim/url-shortener-api/internal/domain/services"
	"github.com/idmaksim/url-shortener-api/internal/infrastructure/services"
	"github.com/labstack/echo/v4"
)

type URLHandler struct {
	urlService domainServices.URLService
}

func NewURLHandler() *URLHandler {
	return &URLHandler{
		urlService: services.NewURLService(),
	}
}

func (h *URLHandler) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (h *URLHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

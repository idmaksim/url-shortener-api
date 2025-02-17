package handlers

import (
	"net/http"

	"github.com/idmaksim/url-shortener-api/internal/domain/services"
	"github.com/labstack/echo/v4"

)

type URLHandler struct {
	urlService *services.URLService
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

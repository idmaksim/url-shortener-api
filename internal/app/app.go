package app

import (
	"fmt"
	"log"

	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/idmaksim/url-shortener-api/internal/delivery/http/handlers"
	"github.com/idmaksim/url-shortener-api/internal/delivery/http/middlewares"
	rateLimiter "github.com/idmaksim/url-shortener-api/internal/delivery/http/rate_limiter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type App struct {
	cfg     *config.Config
	handler *handlers.URLHandler
}

func NewApp() *App {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	return &App{
		cfg:     cfg,
		handler: handlers.NewURLHandler(cfg),
	}
}

func (a *App) Serve() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	a.RegisterRoutes(e)

	return e.Start(fmt.Sprintf(":%d", a.cfg.Http.Port))
}

func (a *App) RegisterRoutes(e *echo.Echo) {
	limiter := rateLimiter.NewIPRateLimiter(10, 10)

	e.POST("/url", a.handler.Create)
	e.GET(":shortURL", a.handler.Get, middlewares.ThrottleMiddleware(limiter))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

package app

import (
	"fmt"
	"log"

	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/idmaksim/url-shortener-api/internal/db"
	"github.com/idmaksim/url-shortener-api/internal/delivery/http/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	db      *db.DB
	cfg     *config.Config
	handler *handlers.URLHandler
}

func NewApp() *App {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := db.NewDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &App{
		db:      db,
		cfg:     cfg,
		handler: handlers.NewURLHandler(),
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
	e.POST("/url", a.handler.Create)
	e.GET("/url/:shortURL", a.handler.Get)
}

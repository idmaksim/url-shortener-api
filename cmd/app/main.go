package main

import (
	"log"

	_ "github.com/idmaksim/url-shortener-api/docs"
	"github.com/idmaksim/url-shortener-api/internal/app"
)

// @title URL Shortener API
// @version 1.0
// @description Service for shortening URLs
// @BasePath /
func main() {
	app := app.NewApp()

	if err := app.Serve(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

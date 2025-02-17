package main

import (
	"log"

	"github.com/idmaksim/url-shortener-api/internal/app"
)

func main() {
	app := app.NewApp()

	if err := app.Serve(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

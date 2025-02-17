package main

import (
	"fmt"
	"log"

	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/idmaksim/url-shortener-api/internal/db"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	database, err := db.NewDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connected successfully:", database.DB)
}

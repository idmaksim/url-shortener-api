package main

import (
	"fmt"
	"log"

	"github.com/idmaksim/url-shortener-api/internal/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Println(*config)
}

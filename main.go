package main

import (
	"go-ecommerce/config"
	"go-ecommerce/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupENV();

	if err != nil {
		log.Fatalf("Config file not loaded properly %v\n", err);
	}

	api.StartServer(cfg);
}
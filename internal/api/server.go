package api

import (
	"go-ecommerce/config"
	"go-ecommerce/internal/api/rest"
	"go-ecommerce/internal/api/rest/handlers"
	"go-ecommerce/internal/domain"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New();
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{});
	rh := &rest.RestHandler{
		App: app,
		DB: db,
	}

	if err != nil {
		log.Fatalf("Database connection error %v\n", err)
	}

	log.Println("Database connected");

	db.AutoMigrate(&domain.User{})

	setupRoutes(rh);
	app.Listen(config.ServerPort);
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoute(rh);
}
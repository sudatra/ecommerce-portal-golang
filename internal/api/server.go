package api

import (
	"go-ecommerce/config"
	"go-ecommerce/internal/api/rest"
	"go-ecommerce/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New();
	rh := &rest.RestHandler{
		App: app,
	}

	setupRoutes(rh);
	app.Listen(config.ServerPort);
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoute(rh);
}
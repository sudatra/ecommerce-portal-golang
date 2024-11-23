package config

import (
	"errors"
	"os"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn string
}

func SetupENV() (cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load();
	}

	httpPort := os.Getenv("HTTP_PORT");
	Dsn := os.Getenv("DSN");

	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("ENV variable not found");
	}

	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("ENV variable not found");
	}

	return AppConfig{ ServerPort: httpPort, Dsn: Dsn }, nil;
}
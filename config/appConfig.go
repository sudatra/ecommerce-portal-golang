package config

import (
	"errors"
	"os"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn string
	AppSecret string
}

func SetupENV() (cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load();
	}

	httpPort := os.Getenv("HTTP_PORT");
	dsn := os.Getenv("DSN");
	appSecret := os.Getenv("APP_SECRET");

	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("ENV variable not found");
	}

	if len(dsn) < 1 {
		return AppConfig{}, errors.New("ENV variable not found");
	}

	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("ENV variable not found");
	}

	return AppConfig{ ServerPort: httpPort, Dsn: dsn, AppSecret: appSecret }, nil;
}
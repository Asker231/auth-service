package config

import (
	"os"

	"github.com/joho/godotenv"
)

type(
	AppConfig struct{
		Auth AuthConfig
		DataBase DataBaseConfig
	}
	AuthConfig struct{
		SECRET string
	}
	DataBaseConfig struct{
		DNS string
	}
)

func LoadAppConfig()*AppConfig{
	godotenv.Load(".env")
	return &AppConfig{
		Auth: AuthConfig{SECRET: os.Getenv("SECRET")},
		DataBase: DataBaseConfig{DNS: os.Getenv("DNS")},
	}
}
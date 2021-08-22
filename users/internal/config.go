package internal

import "os"

type AppConfig struct {
	Port    string
	BaseUrl string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Port:    os.Getenv("PORT"),
		BaseUrl: os.Getenv("BASE_URL"),
	}
}

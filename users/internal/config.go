package internal

import "os"

type AppConfig struct {
	Port string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{Port: os.Getenv("PORT")}
}

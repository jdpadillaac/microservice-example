package internal

import "os"

type AppConfig struct {
	Port        string
	BaseUrl     string
	MongoCnn    string
	MongoDbName string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Port:        os.Getenv("PORT"),
		BaseUrl:     os.Getenv("BASE_URL"),
		MongoCnn:    os.Getenv("MONGO_CNN"),
		MongoDbName: os.Getenv("MONGO_DB_NAME"),
	}
}

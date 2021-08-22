package main

import (
	"github.com/jdpadillaac/microservice-example/tree/main/users/internal"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/infraestructure/rest_api"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig := internal.NewAppConfig()

	restapi.Start(appConfig)
}

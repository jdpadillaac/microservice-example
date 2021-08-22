package mdb

import (
	"context"
	"github.com/jdpadillaac/microservice-example/tree/main/users/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	config *internal.AppConfig
	db     *mongo.Database
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Start(newConfig *internal.AppConfig) {
	config = newConfig
	opts := options.Client().ApplyURI(config.MongoCnn)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("Error en conexión a la base de datos mongodb", err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("error en pin a la base de datos", err.Error())
	}

	db = client.Database(config.MongoDbName)
	log.Println("Base de datos mongo db conectada y lista para usarse")
}

func (c *Config) GeCollection(name string) *mongo.Collection {
	return db.Collection(name)
}

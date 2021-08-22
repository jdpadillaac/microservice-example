package mdb

import (
	"context"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/domain/entity"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/domain/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userDbModel struct {
	ID       string `bson:"_id,omitempty"`
	Names    string `bson:"names,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	RolCode  string `bson:"rol_code,omitempty"`
}

type userRepo struct {
	collection *mongo.Collection
}

func NewUserRepo() repository.User {
	return &userRepo{
		collection: NewConfig().GeCollection("user"),
	}
}

func (u userRepo) Save(user entity.User) (string, error) {

	newUser := userDbModel{
		ID:       user.ID,
		Names:    user.Names,
		Email:    user.Email,
		Password: user.Password,
		RolCode:  user.RolCode,
	}

	result, err := u.collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil

}

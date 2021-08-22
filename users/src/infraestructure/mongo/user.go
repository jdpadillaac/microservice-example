package mdb

import (
	"context"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/domain/entity"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/domain/repository"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
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

func (u userRepo) FindByID(ID string) (entity.User, error) {
	var user entity.User
	var model userDbModel

	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return user, errors.Wrap(err, "Id enviado no válido")
	}

	filter := bson.M{"_id": id}
	err = u.collection.FindOne(context.Background(), filter).Decode(&model)
	if err != nil {
		return user, err
	}

	user = entity.User{
		ID:       model.ID,
		Names:    model.Names,
		Email:    model.Email,
		Password: model.Password,
		RolCode:  model.RolCode,
	}

	return user, nil
}

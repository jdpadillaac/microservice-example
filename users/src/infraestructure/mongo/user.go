package mdb

import (
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/domain/entity"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/domain/repository"
)

type UserRepo struct{}

func NewUserRepo() repository.User {
	return &UserRepo{}
}

func (u UserRepo) Save(user entity.User) (string, error) {
	panic("implement me")
}

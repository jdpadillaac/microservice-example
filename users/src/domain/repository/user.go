package repository

import "github.com/jdpadillaac/microservice-example/tree/main/users/src/domain/entity"

type User interface {
	Save(u entity.User) (string, error)
}

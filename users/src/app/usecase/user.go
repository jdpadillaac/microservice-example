package usecase

import (
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/app/models"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/domain/entity"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/domain/repository"
	"github.com/pkg/errors"
)

type UserUc struct {
	repo repository.User
}

func NewUserUc(repo repository.User) *UserUc {
	return &UserUc{repo: repo}
}

func (u UserUc) Save(doc models.RqRegisterUser) error {
	newUser := entity.User{
		Names:    doc.Names,
		Email:    doc.Email,
		Password: doc.Password,
		RolCode:  doc.RolCode,
	}

	_, dbErr := u.repo.Save(newUser)
	if dbErr != nil {
		return errors.Wrap(dbErr, "db error")
	}

	return nil
}

func (u UserUc) FindByID(ID string) (entity.User, error) {
	return u.repo.FindByID(ID)
}

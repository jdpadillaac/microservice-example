package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jdpadillaac/microservice-example/tree/main/users/internal"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/app/models"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/app/usecase"
	"github.com/jdpadillaac/microservice-example/tree/main/users/src/infraestructure/mongo"
	"net/http"
)

func UserHandler(r *gin.Engine, config *internal.AppConfig) {
	endpoint := config.BaseUrl + "user"
	useCase := usecase.NewUserUc(mdb.NewUserRepo())
	val := validator.New()

	r.Group(endpoint)
	r.POST("", func(c *gin.Context) {
		var model models.RqRegisterUser

		err := c.ShouldBindJSON(&model)
		err = val.Struct(model)
		if err != nil {
			c.JSON(http.StatusBadRequest, newResponse(err.Error(), false))
			return
		}

		err = useCase.Save(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, newResponse(err.Error(), false))
			return
		}

		c.JSON(http.StatusCreated, newResponse("register ok", true))
	})
}

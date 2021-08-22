package restapi

import (
	"fmt"
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

	fmt.Println(endpoint)
	useCase := usecase.NewUserUc(mdb.NewUserRepo())
	val := validator.New()
	user := r.Group(endpoint)

	//Register user
	user.POST("", func(c *gin.Context) {
		var model models.RqRegisterUser

		err := c.BindJSON(&model)
		if err != nil {
			fmt.Println("a")
			c.JSON(http.StatusBadRequest, newResponse(err.Error(), false))
			return
		}

		err = val.Struct(model)
		if err != nil {
			fmt.Println("b")
			c.JSON(http.StatusBadRequest, newResponse(err.Error(), false))
			return
		}

		err = useCase.Save(model)
		if err != nil {
			fmt.Println("c")
			c.JSON(http.StatusInternalServerError, newResponse(err.Error(), false))
			return
		}

		c.JSON(http.StatusCreated, newResponse("register ok", true))
	})

	//Get user by ID
	user.GET("/:ID", func(c *gin.Context) {
		id := c.Param("ID")

		result, err := useCase.FindByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, newResponse(err.Error(), false))
			return
		}

		c.JSON(http.StatusOK, newResponse(result, true))
	})
}

package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/jdpadillaac/microservice-example/tree/main/users/internal"
)

func UserHandler(r *gin.Engine, config *internal.AppConfig) {
	endpoint := config.BaseUrl + "user"

	r.Group(endpoint)

	r.POST("")
}

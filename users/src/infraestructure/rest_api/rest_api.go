package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/jdpadillaac/microservice-example/tree/main/users/internal"
)

func Start(config *internal.AppConfig) {

	r := gin.Default()

	_ = r.Run(":" + config.Port)
}

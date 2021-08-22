package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/jdpadillaac/microservice-example/tree/main/users/internal"
)

type response struct {
	Success bool
	Data    interface{}
}

func newResponse(data interface{}, success bool) *response {
	return &response{
		Success: success,
		Data:    data,
	}
}

func Start(config *internal.AppConfig) {

	r := gin.Default()

	UserHandler(r, config)

	_ = r.Run(":" + config.Port)
}

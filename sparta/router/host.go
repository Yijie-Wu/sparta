package router

import (
	"github.com/gin-gonic/gin"
	"sparta/api"
)

func InitHostRoutes() {
	RegisterRoute(func(rgPublic, rgAuth *gin.RouterGroup) {
		hostAPI := api.NewHostAPI()

		rgAuthHost := rgAuth.Group("host") // 创建host子组

		{
			rgAuthHost.POST("/shutdown", hostAPI.Shutdown)
		}
	})
}

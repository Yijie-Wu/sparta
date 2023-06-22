package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sparta/api"
)

func InitUserRoutes() {
	RegisterRoute(func(rgPublic, rgAuth *gin.RouterGroup) {
		userAPI := api.NewUserAPI()
		rgPublicUser := rgPublic.Group("user")

		{
			rgPublicUser.POST("/login", userAPI.Login)
		}

		rgAuthUser := rgAuth.Group("user") // 创建user子组

		{
			rgAuthUser.GET("", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "yijie"},
						{"id": 2, "name": "hongping"},
					},
				})
			})
			rgAuthUser.GET("/:id", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"id":   1,
					"name": "yijie",
				})
			})
		}
	})
}

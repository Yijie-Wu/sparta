package router

import (
	"github.com/gin-gonic/gin"
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
			rgAuthUser.POST("", userAPI.AddUser)
			rgAuthUser.GET("/:id", userAPI.GetUserByID)
			rgAuthUser.PUT("/:id", userAPI.UpdateUser)
			rgAuthUser.DELETE("/:id", userAPI.DeleteUserByID)
			rgAuthUser.POST("/list", userAPI.GetUserList)
		}
	})
}

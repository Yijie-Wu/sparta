package api

import (
	"github.com/gin-gonic/gin"
	"sparta/service/dto"
)

type UserAPI struct {
	BaseAPI
}

func NewUserAPI() UserAPI {
	return UserAPI{
		BaseAPI: NewBaseAPI(),
	}
}

// @Tags 用户管理
// @Summary 用户登陆
// @Description 登陆系统
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登陆成功"
// @Failure 401 {string} string "登陆失败"
// @Router /api/v1/public/user/login [post]
func (u UserAPI) Login(ctx *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	u.OK(ResponseJson{
		Data: iUserLoginDTO,
	})
}

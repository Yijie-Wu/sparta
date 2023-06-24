package api

import (
	"github.com/gin-gonic/gin"
	"sparta/service"
	"sparta/service/dto"
	"sparta/utils"
)

type UserAPI struct {
	BaseAPI
	Service *service.UserService
}

func NewUserAPI() UserAPI {
	return UserAPI{
		BaseAPI: NewBaseAPI(),
		Service: service.NewUserService(),
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

	iUser, err := u.Service.Login(iUserLoginDTO)
	if err != nil {
		u.ClientFail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(iUser.ID, iUser.NT)

	u.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})
}

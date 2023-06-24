package api

import (
	"github.com/gin-gonic/gin"
	"sparta/service"
	"sparta/service/dto"
)

type HostAPI struct {
	BaseAPI
	Service *service.HostService
}

func NewHostAPI() HostAPI {
	return HostAPI{
		Service: service.NewHostService(),
	}
}

// @Tags 主机管理
// @Summary 关闭主机
// @Description 关闭指定主机
// @Param hostIP formData string true "	主机ip"
// @Success 200 {string} string "登陆成功"
// @Failure 401 {string} string "认证失败"
// @Failure 403 {string} string "权限不足"
// @Router /api/v1/host/shutdown [post]
func (a HostAPI) Shutdown(ctx *gin.Context) {
	var iShutdownHostDTO dto.ShutdownHostDTO

	if err := a.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: iShutdownHostDTO}); err != nil {
		return
	}

	if err := a.Service.Shutdown(iShutdownHostDTO); err != nil {
		a.ClientFail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	a.OK(ResponseJson{
		Msg: "Shutdown Success",
	})
}

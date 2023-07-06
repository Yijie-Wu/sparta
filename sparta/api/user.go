package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sparta/service"
	"sparta/service/dto"
)

const (
	LOGIN_USER_ERROR_CODE  = 10001
	ADD_USER_ERROR_CODE    = 10002
	GET_USER_ERROR_CODE    = 10003
	GET_USERS_ERROR_CODE   = 10004
	UPDATE_USER_ERROR_CODE = 10005
	DELETE_USER_ERROR_CODE = 10006
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
// @Failure 500 {string} string "服务端错误"
// @Router /api/v1/public/user/login [post]
func (u UserAPI) Login(ctx *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	iUser, token, err := u.Service.Login(iUserLoginDTO)

	if err == nil {
		err = service.SetLoginUserTokenToRedis(iUser.ID, token)
	}

	if err != nil {
		u.ClientFail(ResponseJson{
			Status: http.StatusUnauthorized,
			Code:   LOGIN_USER_ERROR_CODE,
			Msg:    err.Error(),
		})
		return
	}

	u.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})
}

// @Tags 用户管理
// @Summary 添加用户
// @Description 添加用户
// @Param nt formData string true "NT"
// @Param name formData string true "用户名"
// @Param email formData string true "邮箱地址"
// @Param password formData string true "密码"
// // @Param file formData file true "头像"
// @Success 201 {string} string "添加成功"
// @Failure 401 {string} string "没有登陆"
// @Failure 403 {string} string "权限不足"
// @Failure 500 {string} string "服务端错误"
// @Router /api/v1/user [post]
func (u UserAPI) AddUser(ctx *gin.Context) {
	var iUserDTO dto.UserAddDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iUserDTO}).GetError(); err != nil {
		return
	}

	//file, _ := ctx.FormFile("file")
	//filePath := fmt.Sprintf("./upload/%s", file.Filename)
	//_ = ctx.SaveUploadedFile(file, filePath)
	//iUserDTO.Avatar = filePath

	if err := u.Service.AddUser(&iUserDTO); err != nil {
		u.ServerFail(ResponseJson{
			Code: ADD_USER_ERROR_CODE,
			Msg:  err.Error(),
		})
		return
	}

	u.OK(ResponseJson{
		Status: http.StatusCreated,
		Data:   iUserDTO,
	})
}

// @Tags 用户管理
// @Summary 获取用户
// @Description 获取用户
// @Param id path int true "ID"
// @Success 200 {string} string "获取成功"
// @Failure 401 {string} string "没有登陆"
// @Failure 403 {string} string "权限不足"
// @Failure 500 {string} string "服务端错误"
// @Router /api/v1/user/{id} [get]
func (u UserAPI) GetUserByID(ctx *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iCommonIDDTO, BindURI: true}).GetError(); err != nil {
		return
	}

	iUser, err := u.Service.GetUserByID(&iCommonIDDTO)

	if err != nil {
		u.ServerFail(ResponseJson{
			Code: GET_USER_ERROR_CODE,
			Msg:  err.Error(),
		})
		return
	}
	u.OK(ResponseJson{
		Data: iUser,
	})
}

// @Tags 用户管理
// @Summary 获取用户列表
// @Description 获取用户列表
// @Param page formData int true "Page"
// @Param limit formData int true "Limit"
// @Success 200 {string} string "获取成功"
// @Failure 401 {string} string "没有登陆"
// @Failure 403 {string} string "权限不足"
// @Failure 500 {string} string "服务端错误"
// @Router /api/v1/user/list [post]
func (u UserAPI) GetUserList(ctx *gin.Context) {
	var iUserListDTO dto.UserListDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iUserListDTO}).GetError(); err != nil {
		return
	}

	iUserList, nTotal, err := u.Service.GetUserList(&iUserListDTO)
	if err != nil {
		u.ServerFail(ResponseJson{
			Code: GET_USERS_ERROR_CODE,
			Msg:  err.Error(),
		})
		return
	}

	u.OK(ResponseJson{
		Data:  iUserList,
		Total: nTotal,
	})
}

// @Tags 用户管理
// @Summary 更新用户
// @Description 更新用户
// @Param id path int true "ID"
// @Param nt formData string false "NT"
// @Param name formData string false "用户名"
// @Param email formData string false "邮箱地址"
// @Success 200 {string} string "更新成功"
// @Failure 401 {string} string "没有登陆"
// @Failure 403 {string} string "权限不足"
// @Failure 500 {string} string "服务端错误"
// @Router /api/v1/user/{id} [put]
func (u UserAPI) UpdateUser(ctx *gin.Context) {
	var iUserUpdateDTO dto.UserUpdateDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iUserUpdateDTO, BindAll: true}).GetError(); err != nil {
		return
	}

	if err := u.Service.UpdateUser(&iUserUpdateDTO); err != nil {
		u.ServerFail(ResponseJson{
			Code: UPDATE_USER_ERROR_CODE,
			Msg:  err.Error(),
		})
		return
	}

	u.OK(ResponseJson{})
}

// @Tags 用户管理
// @Summary 删除用户
// @Description 删除用户
// @Param id path int true "ID"
// @Success 200 {string} string "删除成功"
// @Failure 401 {string} string "没有登陆"
// @Failure 403 {string} string "权限不足"
// @Failure 500 {string} string "服务端错误"
// @Router /api/v1/user/{id} [delete]
func (u UserAPI) DeleteUserByID(ctx *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iCommonIDDTO, BindURI: true}).GetError(); err != nil {
		return
	}

	err := u.Service.DeleteUserByID(&iCommonIDDTO)

	if err != nil {
		u.ServerFail(ResponseJson{
			Code: DELETE_USER_ERROR_CODE,
			Msg:  err.Error(),
		})
		return
	}
	u.OK(ResponseJson{})
}

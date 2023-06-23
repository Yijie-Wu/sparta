package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"sparta/service/dto"
	"sparta/utils"
)

type UserAPI struct {
}

func NewUserAPI() UserAPI {
	return UserAPI{}
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

	err := ctx.ShouldBind(&iUserLoginDTO)
	if err != nil {
		ClientFail(ctx, ResponseJson{
			Status: http.StatusUnprocessableEntity,
			Msg:    ParseValidatorErrors(err.(validator.ValidationErrors), &iUserLoginDTO).Error(),
		})
		return
	}

	OK(ctx, ResponseJson{
		Data: iUserLoginDTO,
	})
}

func ParseValidatorErrors(errs validator.ValidationErrors, target any) error {
	var errResult error

	// 通过反射获取指定元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errs {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}

		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}

		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}

	return errResult
}

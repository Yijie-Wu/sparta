package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
	"sparta/global"
	"sparta/utils"
)

type BaseAPI struct {
	Errors error
	Ctx    *gin.Context
	Logger *zap.SugaredLogger
}

type BuildRequestOption struct {
	Ctx               *gin.Context
	DTO               any
	BindParamsFromURI bool
}

func NewBaseAPI() BaseAPI {
	return BaseAPI{
		Logger: global.Logger,
	}
}

func (a *BaseAPI) AddError(errNew error) {
	a.Errors = utils.AppendError(a.Errors, errNew)
}

func (a *BaseAPI) GetError() error {
	return a.Errors
}

func (a *BaseAPI) BuildRequest(option BuildRequestOption) *BaseAPI {
	var errResult error

	// 绑定请求上下文
	a.Ctx = option.Ctx
	// 绑定请求数据
	if option.DTO != nil {
		if option.BindParamsFromURI {
			errResult = a.Ctx.ShouldBindUri(option.DTO)
		} else {
			errResult = a.Ctx.ShouldBind(option.DTO)
		}
	}

	if errResult != nil {
		errResult = a.ParseValidatorErrors(errResult, option.DTO)
		a.AddError(errResult)
		a.ClientFail(ResponseJson{
			Msg: a.GetError().Error(),
		})
	}

	return a
}

func (a *BaseAPI) ParseValidatorErrors(errs error, target any) error {
	var errResult error

	errsValidation, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}

	// 通过反射获取指定元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errsValidation {
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

func (a *BaseAPI) OK(resp ResponseJson) {
	OK(a.Ctx, resp)
}

func (a *BaseAPI) ClientFail(resp ResponseJson) {
	ClientFail(a.Ctx, resp)
}

func (a *BaseAPI) ServerFail(resp ResponseJson) {
	ServerFail(a.Ctx, resp)
}

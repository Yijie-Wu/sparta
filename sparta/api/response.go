package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  int64  `json:"total,omitempty"`
}

func (m ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(m, ResponseJson{})
}

func buildStatus(resp ResponseJson, nDefaultStatus int) int {
	if nDefaultStatus != 0 {
		return nDefaultStatus
	}
	return resp.Status
}

func HTTPResponse(ctx *gin.Context, status int, resp ResponseJson) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
	}

	ctx.AbortWithStatusJSON(status, resp)
}

func OK(ctx *gin.Context, resp ResponseJson) {
	HTTPResponse(ctx, buildStatus(resp, http.StatusOK), resp)
}

func ClientFail(ctx *gin.Context, resp ResponseJson) {
	HTTPResponse(ctx, buildStatus(resp, http.StatusBadRequest), resp)
}

func ServerFail(ctx *gin.Context, resp ResponseJson) {
	HTTPResponse(ctx, buildStatus(resp, http.StatusInternalServerError), resp)
}

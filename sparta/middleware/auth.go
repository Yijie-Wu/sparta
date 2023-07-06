package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sparta/api"
	"sparta/dao"
	"sparta/global"
	"sparta/global/constants"
	"sparta/model"
	"sparta/service"
	"sparta/utils"
	"strconv"
	"strings"
	"time"
)

const (
	TOKEN_NAME               = "Authorization"
	TOKEN_PREFIX             = "Bearer: "
	ERROR_CODE_INVALID_TOKEN = 10401
	RENEW_TOKEN_DURATION     = 10 * 60 * time.Second
)

func tokenErr(c *gin.Context) {
	api.ClientFail(c, api.ResponseJson{
		Status: http.StatusUnauthorized,
		Code:   ERROR_CODE_INVALID_TOKEN,
		Msg:    "Invalid token",
	})
}

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader(TOKEN_NAME)

		// token 不存在 直接返回
		if token == "" || !strings.HasPrefix(token, TOKEN_PREFIX) {
			tokenErr(c)
			return
		}

		// token 无法解析 直接返回
		token = token[len(TOKEN_PREFIX):]
		iJwtCustomClaims, err := utils.ParseToken(token)
		uUserID := iJwtCustomClaims.ID
		if err != nil || uUserID == 0 {
			tokenErr(c)
			return
		}

		redisKey := constants.LOGIN_USER_TOKEN_REDIS_KEY + strconv.Itoa(int(uUserID))

		// 判断token与访问者登陆的token是否一致 不一致直接返回
		stRedisToken, err := global.RedisClient.Get(redisKey)
		if err != nil || token != stRedisToken {
			tokenErr(c)
			return
		}

		// 判断redis里token是否过期 过期直接返回
		tokenExpireDuration, err := global.RedisClient.GetExpireDuration(redisKey)
		if err != nil || tokenExpireDuration <= 0 {
			tokenErr(c)
			return
		}

		// token 续期
		if tokenExpireDuration.Seconds() < RENEW_TOKEN_DURATION.Seconds() {
			newToken, err := service.GenerateAndCacheLoginUserToken(uUserID, iJwtCustomClaims.Name)
			if err != nil {
				tokenErr(c)
				return
			}

			c.Header("token", newToken)
		}
		iUser, err := dao.NewUserDao().GetUserByID(uUserID)
		if err != nil {
			tokenErr(c)
			return
		}

		//c.Set(constants.LOGIN_USER, iUser)

		c.Set(constants.LOGIN_USER, model.LoginUser{
			ID:     iUser.ID,
			NT:     iUser.NT,
			Name:   iUser.Name,
			Email:  iUser.Email,
			Avatar: iUser.Avatar,
		})

		c.Next()
	}
}

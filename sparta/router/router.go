package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os/signal"
	_ "sparta/docs"
	"sparta/global"
	"sparta/middleware"
	"strings"
	"syscall"
	"time"
)

// IFnRegisterRoute 定义一个路由初始化函数
type IFnRegisterRoute = func(rgPublic, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegisterRoute
)

// RegisterRoute 定义一个注册子路由的函数
func RegisterRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

// InitRouters 初始化所有路由
func InitRouters() {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	r := gin.Default()

	// 加载cors中间件
	r.Use(middleware.Cors())

	// 获取并设置public api router 和 auth api router 的prefix
	authApiPrefix := viper.GetString("app.authApiPrefix")
	publicApiPrefix := viper.GetString("app.publicApiPrefix")

	if authApiPrefix == "" {
		authApiPrefix = "/api/v1"
	}

	if publicApiPrefix == "" {
		publicApiPrefix = "/api/v1/public"
	}

	rgPublic := r.Group(publicApiPrefix)
	rgAuth := r.Group(authApiPrefix)

	// 初始化基础平台路由
	initBasePlatformRoutes()

	// 注册自定义校验器
	registerCustomValidator()

	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}

	// 集成swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := viper.GetString("app.port")
	server := viper.GetString("app.server")
	if port == "" {
		port = "8000"
	}

	app := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", server, port),
		Handler: r,
	}
	// 在协程中启动web服务
	go func() {
		global.Logger.Infof("Start sever at: [%s:%s]", server, port)
		if err := app.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Errorf("Start server failed at:%s", err.Error())
			return
		}
	}()

	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := app.Shutdown(ctx); err != nil {
		global.Logger.Errorf("Stop server failed at:%s", err.Error())
		return
	}
	global.Logger.Info("Stop server success")
}

func initBasePlatformRoutes() {
	InitUserRoutes()
	InitHostRoutes()
}

// 注册自定义验证器
func registerCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("my_validator", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && 0 == strings.Index(value, "a") {
					return true
				}
			}
			return false
		}); err != nil {
			return
		}
	}
}

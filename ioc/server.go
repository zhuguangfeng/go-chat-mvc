package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/activity"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/review"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/user"
	"github.com/zhuguangfeng/go-chat/internal/middleware"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"strings"
	"time"
)

func InitWebServer(
	userHandler *user.UserHandler,
	activityHandler *activity.ActivityHandler,
	reviewHandler *review.ReviewHandler,
	mdls []gin.HandlerFunc,
) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	userHandler.RegisterRouter(server)
	activityHandler.RegisterRouter(server)
	reviewHandler.RegisterRouter(server)

	return server
}

func InitGlobalMiddleware(logger logger.Logger, hdl iJwt.JwtHandler) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		//跨域配置
		cors.New(cors.Config{
			//AllowOrigins: []string{"http://localhost:3000"}, //枚举允许那些跨域请求
			AllowHeaders:  []string{"Content-Type", "Authorization"}, //允许的请求头
			ExposeHeaders: []string{"x-jwt-token"},                   //允许前端访问你的后端响应中带的头部
			AllowOriginFunc: func(origin string) bool { //请求地址如果包含localhost可以请求
				return strings.Contains(origin, "localhost")
			},
			MaxAge: time.Hour * 12,
		}),
		middleware.NewLoginJwtMiddlewareBuilder(logger, hdl).CheckLogin(),
	}
}

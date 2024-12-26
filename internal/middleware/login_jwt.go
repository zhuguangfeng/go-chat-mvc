package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zhuguangfeng/go-chat/internal/common"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"net/http"
)

type LoginJwtMiddlewareBuilder struct {
	iJwt.JwtHandler
	logger logger.Logger
}

func NewLoginJwtMiddlewareBuilder(logger logger.Logger, hdl iJwt.JwtHandler) *LoginJwtMiddlewareBuilder {
	return &LoginJwtMiddlewareBuilder{
		logger:     logger,
		JwtHandler: hdl,
	}
}

func (m *LoginJwtMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == common.GoChatServicePath+"/user/pwd-login" ||
			path == common.GoChatServicePath+"/user/register" {
			return
		}

		tokenStr := m.ExtractToken(ctx)
		var uc iJwt.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return iJwt.JwtKey, nil
		})

		if err != nil {
			m.logger.Info("[middleware.checkLogin]", logger.Error(err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token == nil || !token.Valid {
			m.logger.Info("[middleware.checkLogin]非法token", logger.Any("token", token))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		err = m.CheckSession(ctx, uc.Ssid)
		if err != nil {
			m.logger.Info("[middleware.checkLogin]无效token", logger.Any("token", token))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("user", uc)
	}

}

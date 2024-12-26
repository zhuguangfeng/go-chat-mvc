package user

import (
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/common"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/pkg/ginx"
)

type UserHandler struct {
	iJwt.JwtHandler
	svc user.UserService
}

func NewUserHandler(jwtHandler iJwt.JwtHandler, svc user.UserService) *UserHandler {
	return &UserHandler{
		JwtHandler: jwtHandler,
		svc:        svc,
	}
}

func (hdl *UserHandler) RegisterRouter(router *gin.Engine) {
	userG := router.Group(common.GoChatServicePath + "/user")
	{
		userG.POST("/pwd-login", ginx.WrapBody[dtov1.PasswordLoginReq](hdl.PasswordLogin))
		userG.POST("register", ginx.WrapBody[dtov1.CreateUserReq](hdl.UserRegister))
		userG.GET("/user-info", hdl.GetUserInfo)
	}
}

func (hdl *UserHandler) GetUser(ctx *gin.Context) iJwt.UserClaims {
	u, _ := ctx.Get("user")
	return u.(iJwt.UserClaims)
}

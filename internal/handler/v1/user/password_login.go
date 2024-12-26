package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/common"
	userSvc "github.com/zhuguangfeng/go-chat/internal/service/user"
)

func (hdl *UserHandler) PasswordLogin(ctx *gin.Context, req dtov1.PasswordLoginReq) {
	user, err := hdl.svc.UserPwdLogin(ctx, req.Phone, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, userSvc.ErrUserNotFound):
			common.InternalError(ctx, common.UserPhoneNotFount)
			return
		case errors.Is(err, userSvc.ErrInvalidPassword):
			common.InternalError(ctx, common.UserPasswordInvalid)
			return
		default:
			common.InternalError(ctx, common.SystemError)
			return
		}
	}

	err = hdl.SetLoginToken(ctx, user.ID)
	if err != nil {
		common.InternalError(ctx, common.SystemError)
		return
	}

	common.Success(ctx, nil)
}

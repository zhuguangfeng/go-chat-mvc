package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/service/user"
)

func (hdl *UserHandler) UserRegister(ctx *gin.Context, req dtov1.CreateUserReq) {
	err := hdl.svc.UserRegister(ctx, req)
	if err == nil {
		common.Success(ctx, nil)
		return
	}

	if errors.Is(err, user.ErrUserExists) {
		common.InternalError(ctx, common.UserPhoneExists)
		return
	}

	common.Success(ctx, nil)
}

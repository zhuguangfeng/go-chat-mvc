package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"strconv"
)

func (hdl *UserHandler) GetUserInfo(ctx *gin.Context) {
	uid, _ := strconv.ParseInt(ctx.Query("userId"), 10, 64)
	if uid <= 0 {
		uid = hdl.GetUser(ctx).Uid
	}

	user, err := hdl.svc.GetUserInfo(ctx, uid)
	if err != nil {
		common.InternalError(ctx, common.SystemError)
		return
	}

	common.Success(ctx, user.ToDto())
}

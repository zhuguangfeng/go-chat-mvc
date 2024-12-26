package activity

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/internal/service/activity"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

func (hdl *ActivityHandler) CreateActivity(ctx *gin.Context, req dto.CreateActivityReq, uc jwt.UserClaims) {
	req.SponsorID = uc.Uid
	err := hdl.activitySvc.CreateActivity(ctx, req)
	if err != nil {
		if errors.Is(err, activity.ErrTimeFormatFailed) {
			common.InternalError(ctx, common.BadRequestInvalid)
			return
		}
		hdl.logger.Error("创建活动失败", logger.Error(err))
		common.InternalError(ctx, common.SystemError)
		return
	}

	common.Success(ctx, nil)
}

package activity

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	activitySvc "github.com/zhuguangfeng/go-chat/internal/service/activity"
	userSvc "github.com/zhuguangfeng/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"strconv"
)

func (hdl *ActivityHandler) ActivityDetail(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		common.InternalError(ctx, common.BadRequestInvalid)
		return
	}

	//获取活动信息
	activity, err := hdl.activitySvc.GetActivity(ctx, id)
	if err != nil {
		if errors.Is(err, activitySvc.ErrActivityNotFound) {
			common.InternalError(ctx, common.ActivityNotFound)
			return
		}
		hdl.logger.Error("获取活动失败", logger.Int64("activityId", id), logger.Error(err))
		common.InternalError(ctx, common.SystemError)
		return
	}

	//获取用户信息
	user, err := hdl.userSvc.GetUserInfo(ctx, activity.SponsorID)
	if err != nil {
		if errors.Is(err, userSvc.ErrUserNotFound) {
			common.InternalError(ctx, common.UserDeregister)
			return
		}
		hdl.logger.Error("获取用户信息失败", logger.Int64("userId", activity.SponsorID), logger.Error(err))
		common.InternalError(ctx, common.SystemError)
		return
	}

	res := activity.ToDto()
	res.Sponsor = user.ToDto()
	common.Success(ctx, res)
}

package review

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	reviewSvc "github.com/zhuguangfeng/go-chat/internal/service/review"
	"github.com/zhuguangfeng/go-chat/model"
	"strconv"
)

func (hdl *ReviewHandler) ReviewDetail(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		common.InternalError(ctx, common.BadRequestInvalid)
		return
	}

	review, err := hdl.reviewSvc.GetReview(ctx, id)
	if err != nil {
		if errors.Is(err, reviewSvc.ErrReviewNotFound) {
			common.InternalError(ctx, common.BadRequestInvalid)
		}
		common.InternalError(ctx, common.SystemError)
		return
	}

	user, err := hdl.userSvc.GetUserInfo(ctx, review.ReviewerID)
	if err != nil {
		common.InternalError(ctx, common.SystemError)
		return
	}

	var activity model.Activity

	switch review.Biz {
	case model.ReviewBizActivity:
		activity, err = hdl.activitySvc.GetActivity(ctx, review.BizID)
		if err != nil {
			common.InternalError(ctx, common.SystemError)
			return
		}
	}

	res := review.ToDto()
	res.Reviewer = user.ToDto()
	res.Activity = activity.ToDto()
	common.Success(ctx, res)
}

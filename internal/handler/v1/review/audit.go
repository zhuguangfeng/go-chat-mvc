package review

import (
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"time"
)

// Audit 审核
func (hdl *ReviewHandler) Audit(ctx *gin.Context, req dtov1.ReviewReq, uc jwt.UserClaims) {
	review, err := hdl.reviewSvc.GetReview(ctx, req.ID)
	if err != nil {
		common.InternalError(ctx, common.SystemError)
		return
	}
	if review.Status != model.ReviewStatusPendingReview.Uint() {
		common.InternalError(ctx, common.ReviewedNotAudit)
		return
	}

	review.ReviewerID = uc.Uid
	review.ReviewTime = uint(time.Now().Unix())
	review.Opinion = req.Opinion
	review.Status = req.Status

	switch review.Biz {
	case model.ReviewBizActivity:
		//审核成功
		if req.Status == model.ReviewStatusSuccess.Uint() {
			activity, err := hdl.activitySvc.GetActivity(ctx, review.BizID)
			if err != nil {
				hdl.logger.Error("获取活动信息失败", logger.Int64("activityId", review.BizID), logger.Error(err))
				common.InternalError(ctx, common.SystemError)
				return
			}

			review.SponsorID = activity.SponsorID
			review.Status = model.ReviewStatusSuccess.Uint()

			err = hdl.reviewSvc.AuditActivity(ctx, review, model.Group{
				OwnerID:         activity.SponsorID,
				GroupName:       activity.Title + common.DefaultGroupName,
				MaxPeopleNumber: activity.MaxPeopleNumber,
				Category:        model.GroupCategoryActivity.Uint(),
				Status:          model.GroupStatusNormal.Uint(),
			})
			if err != nil {
				common.InternalError(ctx, common.SystemError)
				return
			}

		}
		//审核失败
		if req.Status == model.ReviewStatusPass.Uint() {
			err := hdl.reviewSvc.AuditActivity(ctx, review, model.Group{})
			if err != nil {
				common.InternalError(ctx, common.SystemError)
				return
			}
		}
	}

	common.Success(ctx, nil)
}

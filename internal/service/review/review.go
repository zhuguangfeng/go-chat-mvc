package review

import (
	"context"
	"errors"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	activityEvent "github.com/zhuguangfeng/go-chat/internal/event/activity"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

type reviewService struct {
	logger        logger.Logger
	reviewDao     dao.ReviewDao
	activityEvent activityEvent.Producer
}

func NewReviewService(logger logger.Logger, reviewDao dao.ReviewDao, activityEvent activityEvent.Producer) ReviewService {
	return &reviewService{
		logger:        logger,
		reviewDao:     reviewDao,
		activityEvent: activityEvent,
	}
}

func (svc *reviewService) ReviewList(ctx context.Context, req dtov1.ReviewListReq) ([]model.Review, int64, error) {
	res, count, err := svc.reviewDao.FindReviews(ctx, req)
	if err != nil {
		svc.logger.Error("从数据库获取审核列表失败", logger.Any("req", req), logger.Error(err))
	}
	return res, count, err

}

// AuditActivity 审核活动
func (svc *reviewService) AuditActivity(ctx context.Context, review model.Review, group model.Group) error {
	var (
		err      error
		activity model.Activity
	)
	switch review.Status {
	case model.ReviewStatusSuccess.Uint():
		activity, err = svc.reviewDao.ReviewActivitySuccess(ctx, review, group)
		if err != nil {
			svc.logger.Error("修改审核状态到数据库失败", logger.Int64("reviewId", review.ID), logger.Error(err))
			return err
		}
		err := svc.activityEvent.ProducerSyncActivityEvent(ctx, activityEvent.ActivityEvent{
			ID:                  activity.ID,
			SponsorID:           activity.SponsorID,
			Title:               activity.Title,
			Desc:                activity.Desc,
			Media:               activity.Media,
			AgeRestrict:         activity.AgeRestrict,
			GenderRestrict:      activity.GenderRestrict,
			CostRestrict:        activity.CostRestrict,
			Visibility:          activity.Visibility,
			MaxPeopleNumber:     activity.MaxPeopleNumber,
			CurrentPeopleNumber: activity.CurrentPeopleNumber,
			Address:             activity.Address,
			Category:            activity.Category,
			StartTime:           activity.StartTime,
			DeadlineTime:        activity.DeadlineTime,
			Status:              activity.Status,
			CreatedTime:         activity.CreatedAt,
			UpdatedTime:         activity.UpdatedAt,
		})
		if err != nil {
			svc.logger.Error("发送同步活动事件失败", logger.Error(err))
		}
	case model.ReviewStatusPass.Uint():
		err = svc.reviewDao.ReviewActivityPass(ctx, review)
		if err != nil {
			svc.logger.Error("修改审核状态到数据库失败", logger.Int64("reviewId", review.ID), logger.Error(err))
			return err
		}
	}

	return nil
}

func (svc *reviewService) GetReview(ctx context.Context, id int64) (model.Review, error) {
	review, err := svc.reviewDao.FindReview(ctx, id)
	if err != nil {
		if errors.Is(err, ErrReviewNotFound) {
			return model.Review{}, ErrReviewNotFound
		}
		svc.logger.Error("从数据库获取审核信息失败", logger.Int64("id", id), logger.Error(err))
		return model.Review{}, err
	}
	return review, nil
}

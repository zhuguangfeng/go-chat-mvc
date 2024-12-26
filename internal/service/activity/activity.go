package activity

import (
	"context"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"

	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

type activityService struct {
	logger        logger.Logger
	activityDao   dao.ActivityDao
	activityEsDao dao.ActivityEsDao
}

func NewActivityService(logger logger.Logger, activityDao dao.ActivityDao, activityEsDao dao.ActivityEsDao) ActivityService {
	return &activityService{
		logger:        logger,
		activityDao:   activityDao,
		activityEsDao: activityEsDao,
	}
}

func (svc *activityService) CreateActivity(ctx context.Context, req dtov1.CreateActivityReq) error {
	startTime, err := req.StartTimeToTimestamp()
	if err != nil {
		svc.logger.Warn("时间转换失败", logger.String("startTime", req.StartTime), logger.Error(err))
		return ErrTimeFormatFailed
	}
	deadlineTime, err := req.DeadlineTimeToTimestamp()
	if err != nil {
		svc.logger.Warn("时间转换失败", logger.String("deadlineTime", req.DeadlineTime), logger.Error(err))
		return ErrTimeFormatFailed
	}

	activity := model.Activity{
		SponsorID:       req.SponsorID,
		Title:           req.Title,
		Desc:            req.Desc,
		Media:           req.Media,
		GenderRestrict:  req.GenderRestrict,
		AgeRestrict:     req.AgeRestrict,
		CostRestrict:    req.CostRestrict,
		Visibility:      req.Visibility,
		MaxPeopleNumber: req.MaxPeopleNumber,
		Address:         req.Address,
		StartTime:       startTime,
		DeadlineTime:    deadlineTime,
		Category:        req.Category,
	}

	activity, err = svc.activityDao.InsertActivity(ctx, activity)
	if err != nil {
		svc.logger.Error("插入活动数据到数据库失败", logger.Error(err))
		return err
	}
	return nil
}

func (svc *activityService) GetActivity(ctx context.Context, id int64) (model.Activity, error) {
	activity, err := svc.activityDao.FindActivity(ctx, id)
	if err != nil {
		svc.logger.Error("从数据库获取活动信息失败", logger.Int64("id", id), logger.Error(err))
	}
	return activity, err
}

func (svc *activityService) GetActivitys(ctx context.Context, req dtov1.ActivitySearchReq) ([]model.Activity, int64, error) {
	activityEss, err := svc.activityEsDao.SearchActivity(ctx, req)
	if err != nil {
		svc.logger.Error("从es获取活动列表失败", logger.Error(err))
		res, count, err := svc.activityDao.FindActivitys(ctx, req)
		if err != nil {
			svc.logger.Error("从数据库获取活动列表失败", logger.Error(err))
			return nil, 0, err
		}
		return res, count, err
	}

	res := slice.Map(activityEss, func(idx int, src model.ActivityEs) model.Activity {
		return src.ToModel()
	})
	return res, 0, nil
}

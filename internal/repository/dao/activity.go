package dao

import (
	"context"
	"errors"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/mysqlx"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
	"gorm.io/gorm"
)

var (
	ErrActivityNotFound = errors.New("activity not found")
)

type ActivityDao interface {
	InsertActivity(ctx context.Context, activity model.Activity) (model.Activity, error)
	FindActivity(ctx context.Context, id int64) (model.Activity, error)
	FindActivitys(ctx context.Context, req dtov1.ActivitySearchReq) ([]model.Activity, int64, error)
}

type GormActivityDao struct {
	db *gorm.DB
}

func NewActivityDao(db *gorm.DB) ActivityDao {
	return &GormActivityDao{db: db}
}

func (dao *GormActivityDao) InsertActivity(ctx context.Context, activity model.Activity) (model.Activity, error) {
	err := dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&activity).Error
		if err != nil {
			return err
		}
		return tx.Create(&model.Review{
			Biz:       model.ReviewBizActivity,
			SponsorID: activity.SponsorID,
			BizID:     activity.ID,
			Status:    model.ReviewStatusPendingReview.Uint(),
		}).Error
	})
	return activity, err
}

func (dao *GormActivityDao) FindActivity(ctx context.Context, id int64) (model.Activity, error) {
	var res model.Activity
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if utils.IsRecordNotFoundError(err) {
		return model.Activity{}, ErrActivityNotFound
	}
	return res, err
}

func (dao *GormActivityDao) FindActivitys(ctx context.Context, req dtov1.ActivitySearchReq) ([]model.Activity, int64, error) {
	var (
		res   = make([]model.Activity, 0)
		count int64
	)
	query := dao.db.WithContext(ctx)

	if req.SearchKey != "" {
		query = query.Where("title like ? or desc like ?", "%"+req.SearchKey+"%", "%"+req.SearchKey+"%")
	}
	if req.AgeRestrict > 0 {
		query = query.Where("age_restrict = ?", req.AgeRestrict)
	}
	if req.GenderRestrict > 0 {
		query = query.Where("gender_restrict = ?", req.GenderRestrict)
	}
	if req.Visibility > 0 {
		query = query.Where("visibility = ?", req.Visibility)
	}
	if req.Category > 0 {
		query = query.Where("category = ?", req.Category)
	}
	if req.Status > 0 {
		query = query.Where("status = ?", req.Status)
	}
	startTime := utils.StringToTimeUnix(req.StartTime)
	if startTime > 0 {
		query = query.Where("start_time >= ?", startTime)
	}
	endTime := utils.StringToTimeUnix(req.StartTime)
	if endTime > 0 {
		query = query.Where("start_time <= ?", endTime)
	}

	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = mysqlx.NewDaoBuilder(query).WithPagination(req.PageNum, req.PageSize).DB.Find(&res).Error

	return res, count, err
}

package dao

import (
	"context"
	"errors"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/mysqlx"
	"gorm.io/gorm"
)

var (
	ErrReviewNotFound = gorm.ErrRecordNotFound
)

type ReviewDao interface {
	FindReviews(ctx context.Context, req dtov1.ReviewListReq) ([]model.Review, int64, error)
	ReviewActivitySuccess(ctx context.Context, review model.Review, group model.Group) (model.Activity, error)
	ReviewActivityPass(ctx context.Context, review model.Review) error
	UpdateReview(ctx context.Context, tx *gorm.DB, review model.Review) error
	FindReview(ctx context.Context, id int64) (model.Review, error)
}
type GormReviewDao struct {
	db *gorm.DB
}

func NewReviewDao1(db *gorm.DB) ReviewDao {
	return &GormReviewDao{
		db: db,
	}
}

func (dao *GormReviewDao) ReviewActivitySuccess(ctx context.Context, review model.Review, group model.Group) (model.Activity, error) {
	var activity model.Activity
	return activity, dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := dao.UpdateReview(ctx, tx, review)
		if err != nil {
			return err
		}
		err = tx.Create(&group).Error
		if err != nil {
			return err
		}
		err = tx.Where("id = ?", review.BizID).Updates(model.Activity{
			GroupID: group.ID,
			Status:  model.ActivityStatusSignup.Uint(),
		}).Error
		if err != nil {
			return err
		}

		return tx.Where("id = ?", review.BizID).First(&activity).Error
	})
}

func (dao *GormReviewDao) ReviewActivityPass(ctx context.Context, review model.Review) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := dao.UpdateReview(ctx, tx, review)
		if err == nil {
			return tx.Model(model.Activity{}).Where("id = ?", review.BizID).Update("status", model.ActivityStatusReviewFailed).Error
		}
		return err
	})
}

func (dao *GormReviewDao) UpdateReview(ctx context.Context, tx *gorm.DB, review model.Review) error {
	if tx != nil {
		return tx.Where("id = ?", review.ID).Updates(&review).Error
	}
	return dao.db.WithContext(ctx).Where("id = ?", review.ID).Updates(&review).Error
}

// FindReviews 获取审核列表 分页
func (dao *GormReviewDao) FindReviews(ctx context.Context, req dtov1.ReviewListReq) ([]model.Review, int64, error) {
	var (
		res   = make([]model.Review, 0)
		count int64
	)
	query := dao.db.WithContext(ctx)
	if req.Biz != "" {
		query = query.Where("biz = ?", req.Biz)
	}
	if req.BizID > 0 {
		query = query.Where("biz_id = ?", req.BizID)
	}
	if req.Status > 0 {
		query = query.Where("status = ?", req.Status)
	}
	count, err := dao.getCount(query)
	if err != nil {
		return nil, 0, err
	}
	err = mysqlx.NewDaoBuilder(query).WithPagination(req.PageNum, req.PageSize).DB.Find(&res).Error
	return res, count, err
}

func (dao *GormReviewDao) FindReview(ctx context.Context, id int64) (model.Review, error) {
	var res model.Review
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if errors.Is(err, ErrReviewNotFound) {
		return res, ErrReviewNotFound
	}
	return res, err
}

func (dao *GormReviewDao) getCount(db *gorm.DB) (int64, error) {
	var res int64
	err := db.Model(model.Review{}).Count(&res).Error
	return res, err
}

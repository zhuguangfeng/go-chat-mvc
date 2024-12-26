package review

import (
	"context"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
)

var (
	ErrReviewNotFound = dao.ErrReviewNotFound
)

type ReviewService interface {
	GetReview(ctx context.Context, id int64) (model.Review, error)
	ReviewList(ctx context.Context, req dtov1.ReviewListReq) ([]model.Review, int64, error)
	AuditActivity(ctx context.Context, review model.Review, group model.Group) error
}

package activity

import (
	"context"
	"errors"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
)

var (
	ErrTimeFormatFailed = errors.New("时间格式有误")
	ErrActivityNotFound = dao.ErrActivityNotFound
)

type ActivityService interface {
	CreateActivity(ctx context.Context, req dtov1.CreateActivityReq) error
	GetActivity(ctx context.Context, id int64) (model.Activity, error)
	GetActivitys(ctx context.Context, req dtov1.ActivitySearchReq) ([]model.Activity, int64, error)
}

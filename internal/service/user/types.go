package user

import (
	"context"
	"errors"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
)

var (
	ErrUserExists      = dao.ErrUserDuplicate
	ErrUserNotFound    = dao.ErrUserNotFound
	ErrInvalidPassword = errors.New("密码错误")
)

type UserService interface {
	UserRegister(ctx context.Context, user dtov1.CreateUserReq) error
	UserPwdLogin(ctx context.Context, phone, password string) (model.User, error)
	GetUserInfo(ctx context.Context, uid int64) (model.User, error)
}

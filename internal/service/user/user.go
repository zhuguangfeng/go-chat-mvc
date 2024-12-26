package user

import (
	"context"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/repository/cache"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"time"
)

type userService struct {
	logger    logger.Logger
	userDao   dao.UserDao
	userCache cache.UserCache
}

func NewUserService(logger logger.Logger, userDao dao.UserDao, userCache cache.UserCache) UserService {
	return &userService{
		logger:    logger,
		userDao:   userDao,
		userCache: userCache,
	}
}

// UserRegister 用户注册
func (svc *userService) UserRegister(ctx context.Context, req dtov1.CreateUserReq) error {
	_, err := svc.userDao.InsertUser(ctx, model.User{
		Username:        "小白",
		Avatar:          "123",
		BackgroundImage: "123",
		Phone:           req.Phone,
		Password:        req.Password,
		Status:          model.UserStatusNormal.Uint(),
	})
	if err != nil {
		svc.logger.Error("创建用户失败", logger.Error(err))
		return err
	}
	return nil
}

// UserPwdLogin 账号密码登录
func (svc *userService) UserPwdLogin(ctx context.Context, phone, password string) (model.User, error) {
	user, err := svc.userDao.FindUserByPhone(ctx, phone)
	if err != nil {
		return model.User{}, err
	}
	if user.Password != password {
		return model.User{}, ErrInvalidPassword
	}
	return user, nil
}

// GetUserInfo 获取用户信息
func (svc *userService) GetUserInfo(ctx context.Context, uid int64) (model.User, error) {
	//从缓存获取用户信息
	user, err := svc.userCache.GetUser(ctx, uid)
	if err == nil {
		return user, nil
	}

	svc.logger.Warn("从缓存获取用户信息失败", logger.Int64("uid", uid), logger.Error(err))

	user, err = svc.userDao.FindUser(ctx, uid)
	if err != nil {
		svc.logger.Error("从数据库获取用户信息失败", logger.Int64("uid", uid), logger.Error(err))
		return model.User{}, err
	}

	//回写用户信息到缓存
	err = svc.userCache.SetUser(ctx, user, time.Hour*24*30)
	if err != nil {
		svc.logger.Warn("回写用户信息到数据库失败", logger.Error(err))
	}

	return user, nil
}

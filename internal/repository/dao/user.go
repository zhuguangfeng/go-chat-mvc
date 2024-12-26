package dao

import (
	"context"
	"errors"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicate = errors.New("用户已存在")
	ErrUserNotFound  = errors.New("用户不存在")
)

type UserDao interface {
	InsertUser(ctx context.Context, user model.User) (model.User, error)
	FindUser(ctx context.Context, id int64) (model.User, error)
	FindUserByPhone(ctx context.Context, phone string) (model.User, error)
}

type GormUserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &GormUserDao{
		db: db,
	}
}

func (dao *GormUserDao) InsertUser(ctx context.Context, user model.User) (model.User, error) {
	err := dao.db.WithContext(ctx).Create(&user).Error
	if utils.IsDuplicateKeyError(err) {
		return model.User{}, ErrUserDuplicate
	}
	return user, err
}

func (dao *GormUserDao) FindUser(ctx context.Context, id int64) (model.User, error) {
	var res model.User
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if utils.IsRecordNotFoundError(err) {
		return res, ErrUserNotFound
	}
	return res, err
}

func (dao *GormUserDao) FindUserByPhone(ctx context.Context, phone string) (model.User, error) {
	var res model.User
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).First(&res).Error
	if utils.IsRecordNotFoundError(err) {
		return res, ErrUserNotFound
	}
	return res, err
}

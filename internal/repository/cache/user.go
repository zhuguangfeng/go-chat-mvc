package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhuguangfeng/go-chat/model"

	"time"
)

var (
	UserInfoKeyPrefix = "user:userInfo:"
)

type UserCache interface {
	GetUser(ctx context.Context, uid int64) (model.User, error)
	SetUser(ctx context.Context, user model.User, expiration time.Duration) error
}

type RedisUserCache struct {
	redisCli redis.Cmdable
}

func NewUserCache(redisCli redis.Cmdable) UserCache {
	return &RedisUserCache{
		redisCli: redisCli,
	}
}
func (cache *RedisUserCache) GetUser(ctx context.Context, uid int64) (model.User, error) {
	var res model.User
	err := cache.redisCli.Get(ctx, cache.getUserKey(uid)).Scan(&res)
	return res, err
}

func (cache *RedisUserCache) SetUser(ctx context.Context, user model.User, expiration time.Duration) error {
	return cache.redisCli.Set(ctx, cache.getUserKey(user.ID), &user, expiration).Err()
}

func (cache *RedisUserCache) getUserKey(uid int64) string {
	return fmt.Sprintf("%s%d", UserInfoKeyPrefix, uid)
}

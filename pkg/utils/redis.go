package utils

import (
	"errors"
	"github.com/redis/go-redis/v9"
)

func IsRedisNilError(err error) bool {
	return errors.Is(err, redis.Nil)
}

package ioc

import (
	"github.com/spf13/viper"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"go.uber.org/zap"
)

// InitLogger 初始化日志
func InitLogger() logger.Logger {
	cfg := zap.NewDevelopmentConfig()

	err := viper.UnmarshalKey("log", &cfg)
	if err != nil {
		panic(err)
	}

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger.NewZapLogger(l)
}

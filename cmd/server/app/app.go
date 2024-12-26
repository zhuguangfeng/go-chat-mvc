package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/pkg/saramax"
)

type App struct {
	Server    *gin.Engine
	Consumers []saramax.Consumer
}

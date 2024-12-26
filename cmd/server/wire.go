//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zhuguangfeng/go-chat/cmd/server/app"
	activityEvent "github.com/zhuguangfeng/go-chat/internal/event/activity"
	activityHdl "github.com/zhuguangfeng/go-chat/internal/handler/v1/activity"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	reviewHdl "github.com/zhuguangfeng/go-chat/internal/handler/v1/review"
	userHdl "github.com/zhuguangfeng/go-chat/internal/handler/v1/user"
	"github.com/zhuguangfeng/go-chat/internal/repository/cache"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/internal/service/activity"
	"github.com/zhuguangfeng/go-chat/internal/service/review"
	"github.com/zhuguangfeng/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/ioc"
)

func InitWebServer() *app.App {
	wire.Build(
		ioc.InitLogger,
		ioc.InitMysql,
		ioc.InitRedisCmd,
		ioc.InitGlobalMiddleware,
		ioc.InitWebServer,
		ioc.InitEsClient,
		ioc.InitKafka,
		activityEvent.NewActivityConsumer,
		ioc.InitSaramaSyncProducer,
		ioc.NewConsumers,

		activityEvent.NewProducer,

		dao.NewUserDao,
		cache.NewUserCache,
		dao.NewActivityDao,
		dao.NewReviewDao1,
		dao.NewActivityEsDao,

		user.NewUserService,
		activity.NewActivityService,
		review.NewReviewService,

		iJwt.NewJwtHandler,
		userHdl.NewUserHandler,
		activityHdl.NewActivityHandler,
		reviewHdl.NewReviewHandler,

		wire.Struct(new(app.App), "*"),
	)

	return new(app.App)
}

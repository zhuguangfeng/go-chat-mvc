package review

import (
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/internal/service/activity"
	"github.com/zhuguangfeng/go-chat/internal/service/review"
	"github.com/zhuguangfeng/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/pkg/ginx"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

type ReviewHandler struct {
	logger      logger.Logger
	reviewSvc   review.ReviewService
	activitySvc activity.ActivityService
	userSvc     user.UserService
}

func NewReviewHandler(logger logger.Logger,
	reviewSvc review.ReviewService,
	activitySvc activity.ActivityService,
	userSvc user.UserService,
) *ReviewHandler {
	return &ReviewHandler{
		logger:      logger,
		reviewSvc:   reviewSvc,
		activitySvc: activitySvc,
		userSvc:     userSvc,
	}
}

func (hdl *ReviewHandler) RegisterRouter(router *gin.Engine) {
	path := common.GoChatServicePath + "/review"
	router.POST(path+"/list", ginx.WrapBody[dtov1.ReviewListReq](hdl.ReviewList))
	router.POST(path+"/audit", ginx.WrapBodyAndClaims[dtov1.ReviewReq, jwt.UserClaims](hdl.Audit))

	router.GET(path+"/detail", hdl.ReviewDetail)
}

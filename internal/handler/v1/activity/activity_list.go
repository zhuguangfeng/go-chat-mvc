package activity

import (
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
)

func (hdl *ActivityHandler) ActivityList(ctx *gin.Context, req dtov1.ActivitySearchReq) {
	activitys, count, err := hdl.activitySvc.GetActivitys(ctx, req)
	if err != nil {
		hdl.logger.Error("获取活动列表失败", logger.Error(err))
		common.InternalError(ctx, common.SystemError)
		return
	}

	common.Success(ctx, common.ListObj{
		CurrentCount: len(activitys),
		TotalCount:   count,
		TotalPage:    utils.GetPageCount(int(count), req.PageSize),
		Result: slice.Map(activitys, func(idx int, src model.Activity) dtov1.Activity {
			return src.ToDto()
		}),
	})

}

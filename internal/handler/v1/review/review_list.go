package review

import (
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
)

func (hdl *ReviewHandler) ReviewList(ctx *gin.Context, req dtov1.ReviewListReq) {
	reviews, count, err := hdl.reviewSvc.ReviewList(ctx, req)
	if err != nil {
		common.InternalError(ctx, common.SystemError)
		return
	}

	common.Success(ctx, common.ListObj{
		CurrentCount: len(reviews),
		TotalCount:   count,
		TotalPage:    utils.GetPageCount(int(count), req.PageSize),
		Result: slice.Map(reviews, func(idx int, src model.Review) dtov1.Review {
			return src.ToDto()
		}),
	})

}

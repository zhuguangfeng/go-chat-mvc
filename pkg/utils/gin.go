package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPagination(ctx *gin.Context) (int, int) {
	var (
		pageNum, _  = strconv.Atoi(ctx.Param("pageNum"))
		PageSize, _ = strconv.Atoi(ctx.Param("pageSize"))
	)

	if pageNum <= 0 {
		pageNum = 1
	}
	if PageSize <= 0 {
		PageSize = 10
	}
	return pageNum, PageSize
}

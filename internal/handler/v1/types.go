package v1

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterRouter(router *gin.Engine)
}

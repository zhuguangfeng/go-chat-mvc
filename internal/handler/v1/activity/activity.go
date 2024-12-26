package activity

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"github.com/gorilla/websocket"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	activitySvc "github.com/zhuguangfeng/go-chat/internal/service/activity"
	"github.com/zhuguangfeng/go-chat/internal/service/thirdparty/baidu"
	userSvc "github.com/zhuguangfeng/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/pkg/ginx"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"os"
)

type ActivityHandler struct {
	logger      logger.Logger
	activitySvc activitySvc.ActivityService
	userSvc     userSvc.UserService
}

func NewActivityHandler(logger logger.Logger, activitySvc activitySvc.ActivityService, userSvc userSvc.UserService) *ActivityHandler {
	return &ActivityHandler{
		logger:      logger,
		activitySvc: activitySvc,
		userSvc:     userSvc,
	}
}

func (hdl *ActivityHandler) RegisterRouter(router *gin.Engine) {
	activityG := router.Group(common.GoChatServicePath + "/activity")
	{
		activityG.POST("/create", ginx.WrapBodyAndClaims[dtov1.CreateActivityReq, jwt.UserClaims](hdl.CreateActivity))
		activityG.GET("/detail", hdl.ActivityDetail)
		activityG.POST("/search", ginx.WrapBody[dtov1.ActivitySearchReq](hdl.ActivityList))

		activityG.POST("/ws", hdl.TestBaiDuWs)
	}
}

func (hdl *ActivityHandler) TestBaiDuWs(ctx *gin.Context) {
	uuid := uuid2.New().String()
	conn, _, err := websocket.DefaultDialer.Dial("wss://vop.baidu.com/realtime_asr?sn="+uuid, nil)
	if err != nil {
		common.InternalError(ctx, common.SystemError)
		return
	}
	vgSvc := baidu.NewVoiceRecognition(conn)
	bytes := ReadFile("D:\\GoProject\\src\\zgf\\go-chat\\go-chat-web\\internal\\handler\\v1\\activity\\16k.pcm")
	res, err := vgSvc.Identify(ctx, bytes, 6000)
	if err != nil {
		common.InternalError(ctx, common.SystemError)
		return
	}
	common.Success(ctx, res)
}

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return nil
	}
	defer file.Close()

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Failed to get file info: %v\n", err)
		return nil
	}
	fileSize := fileInfo.Size()

	// 读取PCM数据到切片
	pcmData := make([]byte, fileSize)
	_, err = file.Read(pcmData)
	if err != nil {
		fmt.Printf("Failed to read PCM data: %v\n", err)
		return nil
	}

	return pcmData
}

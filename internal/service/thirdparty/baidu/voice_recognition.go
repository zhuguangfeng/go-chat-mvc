package baidu

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zhuguangfeng/go-chat/pkg/request"
)

type VoiceRecognition struct {
	WebsocketClient *request.WebSocketClient
}

func NewVoiceRecognition(conn *websocket.Conn) *VoiceRecognition {
	return &VoiceRecognition{
		WebsocketClient: request.NewWebSocketClient(conn),
	}
}

type StartReq struct {
	Type string `json:"type"`
	Data struct {
		AppId  int    `json:"appid"`
		AppKey string `json:"appkey"`
		DevPid int    `json:"dev_pid"`
		LmId   int    `json:"lm_id"`
		CuId   string `json:"cuid"`
		Format string `json:"format"`
		Sample int    `json:"sample"`
	} `json:"data"`
}

type Opcode struct {
	Type string `json:"type"`
}

type BaiduResult struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	Type   string `json:"type"`
	Result string `json:"result"`
}

func (v *VoiceRecognition) Identify(ctx context.Context, data []byte, fileLen int) (string, error) {
	startReq := StartReq{
		Type: "START",
		Data: struct {
			AppId  int    `json:"appid"`
			AppKey string `json:"appkey"`
			DevPid int    `json:"dev_pid"`
			LmId   int    `json:"lm_id"`
			CuId   string `json:"cuid"`
			Format string `json:"format"`
			Sample int    `json:"sample"`
		}{
			AppId:  116835869,
			AppKey: "LlO0bh8SUe0ikhm3nOgfNtae",
			DevPid: 15372,
			CuId:   "cuid-1",
			Format: "pcm",
			Sample: 16000,
		},
	}

	err := v.sendStartRequest(startReq)
	if err != nil {
		fmt.Println("发送开始任务失败", err)
		return "", err
	}

	err = v.sendAudioData(data, fileLen)
	if err != nil {
		fmt.Println("发送音频数据失败", err)
		return "", err
	}

	err = v.sendEndRequest(Opcode{
		Type: "FINISH",
	})
	if err != nil {
		fmt.Println("发送结束数据失败", err)
		return "", err
	}

	msg, err := v.ReadMsg()
	if err != nil {
		fmt.Println("读取数据失败", err)
		return "", err
	}

	return msg, nil

}

func (v *VoiceRecognition) sendStartRequest(req StartReq) (err error) {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return v.WebsocketClient.Conn.WriteMessage(websocket.TextMessage, bytes)
}

func (v *VoiceRecognition) sendAudioData(data []byte, chunk_ms int) error {

	index := 0
	total := len(data)
	chunkLen := 16000 * 2 / 1000 * chunk_ms

	for index < total {
		end := index + chunkLen
		if end >= total {
			end = total
		}
		body := data[index:end]
		err := v.WebsocketClient.Conn.WriteMessage(websocket.BinaryMessage, body)
		if err != nil {
			return err
		}

		index = end
	}
	return nil
}

func (v *VoiceRecognition) sendEndRequest(req Opcode) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return v.WebsocketClient.Conn.WriteMessage(websocket.TextMessage, bytes)
}

func (v *VoiceRecognition) sendCloseRequest(req Opcode) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return v.WebsocketClient.Conn.WriteMessage(websocket.TextMessage, bytes)
}

func (v *VoiceRecognition) ReadMsg() (string, error) {
	for {
		typ, msg, err := v.WebsocketClient.Conn.ReadMessage()
		if err != nil {
			return "", err
		}
		fmt.Println(typ)
		fmt.Println(string(msg))

		var res BaiduResult
		err = json.Unmarshal(msg, &res)
		if err != nil {
			return "", err
		}

		if res.ErrNo != 0 {
			return "", errors.New(string(msg))
		}

		if res.Type == "FIN_TEXT" {
			return res.Result, nil
		}

	}
}

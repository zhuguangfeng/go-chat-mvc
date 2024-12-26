package request

import "github.com/gorilla/websocket"

type WebSocketClient struct {
	Conn *websocket.Conn
}

func NewWebSocketClient(conn *websocket.Conn) *WebSocketClient {
	return &WebSocketClient{Conn: conn}
}

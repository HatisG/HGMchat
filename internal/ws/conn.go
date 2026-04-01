package ws

import (
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	UserID   uint
	SendChan chan []byte
	IsClosed bool
}

// 新用户上线
func NewClient(userID uint, conn *websocket.Conn) *Client {
	//设置超时自动下线
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	return &Client{
		UserID:   userID,
		Conn:     conn,
		SendChan: make(chan []byte, 256),
		IsClosed: false,
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Conn.Close()
		Server.Offline(c.UserID)
	}()

	//心跳机制
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {

		//发送消息
		case msg, ok := <-c.SendChan:
			if !ok {
				return
			}
			if c.IsClosed {
				return
			}
			_ = c.Conn.WriteMessage(websocket.TextMessage, msg)

			//发送心跳
		case <-ticker.C:
			if c.IsClosed {
				return
			}
			_ = c.Conn.WriteMessage(websocket.PingMessage, nil)
		}

	}

}

func (c *Client) Read() {
	defer func() {
		_ = c.Conn.Close()
		Server.Offline(c.UserID)
		c.IsClosed = true
	}()

	for {
		//读取消息
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		Server.HandleMessage(c.UserID, msg)
	}

}

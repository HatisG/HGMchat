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

func NewClient(userID uint, conn *websocket.Conn) *Client {

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

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-c.SendChan:
			if !ok {
				return
			}
			_ = c.Conn.WriteMessage(websocket.TextMessage, msg)

		case <-ticker.C:
			_ = c.Conn.WriteMessage(websocket.PingMessage, nil)
		}

	}

}

func (c *Client) Read() {
	defer func() {
		_ = c.Conn.Close()
		Server.Offline(c.UserID)
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		Server.HandleMessage(c.UserID, msg)
	}

}

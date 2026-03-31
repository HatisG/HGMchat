package ws

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	UserID   uint
	SendChan chan []byte
	IsClosed bool
}

func NewClient(userID uint, conn *websocket.Conn) *Client {
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
	}()

	for msg := range c.SendChan {
		_ = c.Conn.WriteMessage(websocket.TextMessage, msg)
	}

}

func (c *Client) Read() {
	defer func() {
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

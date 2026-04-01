package ws

import (
	"HGMchat/internal/dao"
	"HGMchat/internal/model"
	"encoding/json"
	"sync"
	"time"
)

var Server = NewWSServer()

type WSServer struct {
	OnlineMap map[uint]*Client
	sync.RWMutex
	msgChan chan model.Message
}

func NewWSServer() *WSServer {
	ws := &WSServer{
		OnlineMap: make(map[uint]*Client),
		msgChan:   make(chan model.Message, 1024),
	}

	go ws.batchInsertWorker()
	return ws
}

func (s *WSServer) Online(userID uint, client *Client) {

	s.Lock()
	defer s.Unlock()

	s.OnlineMap[userID] = client

}

func (s *WSServer) Offline(userID uint) {
	s.Lock()
	defer s.Unlock()

	if client, ok := s.OnlineMap[userID]; ok {
		close(client.SendChan)

		delete(s.OnlineMap, userID)
	}
}

func (s *WSServer) SendToUser(userID uint, msg []byte) bool {
	s.RLock()
	defer s.RUnlock()

	client, ok := s.OnlineMap[userID]
	if !ok {
		return false
	}

	client.SendChan <- msg
	return true
}

func (s *WSServer) HandleMessage(fromUserID uint, msgData []byte) {

	var msg struct {
		ToUserID uint   `json:"to_user_id"`
		Content  string `json:"content"`
		Type     int    `json:"type"`
	}

	if err := json.Unmarshal(msgData, &msg); err != nil {
		return
	}

	select {
	case s.msgChan <- model.Message{
		FromUserID: fromUserID,
		ToUserID:   msg.ToUserID,
		Content:    msg.Content,
		Type:       msg.Type,
	}:

	default:

	}

	s.SendToUser(msg.ToUserID, msgData)
}

func (s *WSServer) batchInsertWorker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	var msgList []model.Message

	for {
		select {
		case msg := <-s.msgChan:
			msgList = append(msgList, msg)

			if len(msgList) >= 100 {
				_ = dao.BatchCreateMessage(msgList)
				msgList = nil
			}

		case <-ticker.C:
			if len(msgList) > 0 {
				_ = dao.BatchCreateMessage(msgList)
				msgList = nil
			}

		}

	}

}

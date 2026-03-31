package ws

import (
	"encoding/json"
	"sync"
)

var Server = NewWSServer()

type WSServer struct {
	OnlineMap map[uint]*Client
	sync.RWMutex
}

func NewWSServer() *WSServer {
	return &WSServer{
		OnlineMap: make(map[uint]*Client),
	}
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

func (s *WSServer) HandleMessage(fromUserID uint, msg []byte) {

	var massage Message
	if err := json.Unmarshal(msg, &massage); err != nil {
		return
	}

	s.SendToUser(massage.ToUserID, msg)
}

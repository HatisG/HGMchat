package vo

type MessageVO struct {
	ID         uint   `json:"id"`
	FromUserID uint   `json:"from_user_id"`
	ToUserID   uint   `json:"to_user_id"`
	Content    string `json:"content"`
	Type       int    `json:"type"`
	CreateAt   string `json:"create_at"`
}

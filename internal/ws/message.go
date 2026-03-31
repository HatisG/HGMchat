package ws

type Message struct {
	FromUserID uint
	ToUserID   uint
	Content    string
	Type       int
	CreateTime int64
}

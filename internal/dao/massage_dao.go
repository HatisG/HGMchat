package dao

import "HGMchat/internal/model"

//创建消息
func CreateMessage(msg *model.Message) error {
	return DB.Create(msg).Error
}

//
func BatchCreateMessage(msgs []model.Message) error {
	if len(msgs) == 0 {
		return nil
	}
	return DB.Create(&msgs).Error
}

//获取对话历史
func GetChatHistory(fromUID, toUID uint, limit int) ([]model.Message, error) {
	var list []model.Message

	err := DB.Where("from_user_id = ? and to_user_id = ? or from_user_id = ? and to_user_id = ?",
		fromUID, toUID, toUID, fromUID).
		Order("create_at ASC").
		Limit(limit).
		Find(&list).Error

	return list, err
}

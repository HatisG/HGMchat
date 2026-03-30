package dao

import (
	"HGMchat/internal/model"
)

func IsUsernameExists(username string) bool {
	var count int64
	//查找并计数
	DB.Model(&model.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func Create(user *model.User) error {
	return DB.Create(user).Error
}

func GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

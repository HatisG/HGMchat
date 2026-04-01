package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromUserID uint   `gorm:"index;not null"`
	ToUserID   uint   `gorm:"index;not null"`
	Content    string `gorm:"type:text;not null"`
	Type       int    `gorm:"default:1"`
}

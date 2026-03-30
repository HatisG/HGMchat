package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitMySQL() {
	//数据库dsn
	dsn := "root:root@tcp(127.0.0.1:3306)/HGMchat?charset=utf8mb4&parseTime=True&loc=Local"

	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //打印sql
	})

	if err != nil {
		panic("Failed to connect databases:" + err.Error())
	}

	DB = db
}

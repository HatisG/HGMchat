package main

import (
	"HGMchat/internal/dao"
	"HGMchat/internal/model"
)

func main() {

	dao.InitMySQL()

	dao.DB.AutoMigrate(
		&model.User{},
	)

	r := SetupRouter()

	r.Run(":8080")
}

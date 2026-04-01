package main

import (
	"HGMchat/internal/dao"
	"HGMchat/internal/handler"
	"HGMchat/internal/middleware"
	"HGMchat/internal/model"
	"HGMchat/internal/ws"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
	}

	wsGroup := r.Group("/ws")
	wsGroup.Use(middleware.JWT())
	{
		wsGroup.GET("", ws.WSHandler)
	}

	chatGroup := r.Group("/chat")
	chatGroup.Use(middleware.JWT())
	{
		wsGroup.POST("/history", handler.GetChatHistory)
	}
	return r
}

func main() {

	dao.InitMySQL()

	dao.DB.AutoMigrate(
		&model.User{},
	)

	r := SetupRouter()

	r.Run(":8080")
}

package handler

import (
	"HGMchat/internal/dao"
	"HGMchat/internal/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetChatHistory(c *gin.Context) {
	fromUID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "未登录"})
		return
	}

	var req struct {
		ToUserID uint `json:"to_user_id" binding:"required"`
		Limit    int  `json:"limit"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "参数错误"})
		return
	}

	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 50
	}

	list, err := dao.GetChatHistory(fromUID.(uint), req.ToUserID, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "获取失败"})
		return
	}

	var res []vo.MessageVO
	for _, msg := range list {
		res = append(res, vo.MessageVO{
			ID:         msg.ID,
			FromUserID: msg.FromUserID,
			ToUserID:   msg.ToUserID,
			Content:    msg.Content,
			Type:       msg.Type,
			CreateAt:   msg.CreatedAt.Format("2006-01-02 15:04:05"),
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": res,
	})

}

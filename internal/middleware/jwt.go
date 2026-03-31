package middleware

import (
	"HGMchat/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录"})
			c.Abort()
			return
		}

		claims, err := pkg.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}

}

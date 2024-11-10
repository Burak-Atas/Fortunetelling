package middleware

import (
	"github.com/Burak-Atas/kahve_fali/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token != "" {
			claims, msg := jwt.ValidateToken(token)
			if msg != "" {
				c.JSON(400, gin.H{
					"message": msg,
				})
				return
			}
			c.Set("uid", claims.Uid)
			c.Next()
			return
		}
		c.JSON(403, gin.H{
			"message": "yetkisiz erişim",
		})
	}
}

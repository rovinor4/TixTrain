package middleware

import (
	"TixTrain/app/model"
	"TixTrain/database"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		if tokenString == "" {
			c.JSON(401, gin.H{
				"message": "Token tidak ditemukan",
			})
			c.Abort()
			return
		}

		// Cari token di database
		var token model.Token
		dataToken := database.DB.Where("value = ?", tokenString).First(&token)
		if dataToken.RowsAffected == 0 || dataToken.Error != nil {
			c.JSON(401, gin.H{
				"message": "Token tidak valid",
			})
			c.Abort()
			return
		}

		if token.ExpiresAt.Before(time.Now()) {
			c.JSON(401, gin.H{
				"message": "Token sudah kadaluarsa",
			})
			c.Abort()
			return
		}

		var user model.User
		dataUser := database.DB.Where("id = ?", token.UserID).First(&user)
		if dataUser.RowsAffected == 0 || dataUser.Error != nil {
			c.JSON(401, gin.H{
				"message": "User tidak ditemukan",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("user_id", user.ID)
		c.Set("user_role", user.Role)
		c.Set("token", tokenString)

		c.Next()
	}
}

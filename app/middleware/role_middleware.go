package middleware

import "github.com/gin-gonic/gin"

func RoleMiddleware(Role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists || userRole != Role {
			c.JSON(403, gin.H{
				"message": "Akses ditolak",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

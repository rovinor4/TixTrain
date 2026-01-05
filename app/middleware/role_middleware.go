package middleware

import "github.com/gin-gonic/gin"

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoles, exists := c.Get("user_roles")
		if !exists {
			c.JSON(403, gin.H{
				"message": "Akses ditolak",
			})
			c.Abort()
			return
		}

		// userRoles is a []string
		roles, ok := userRoles.([]string)
		if !ok {
			c.JSON(403, gin.H{
				"message": "Akses ditolak",
			})
			c.Abort()
			return
		}

		// Check if user has the required role
		hasRole := false
		for _, role := range roles {
			if role == requiredRole {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(403, gin.H{
				"message": "Akses ditolak",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

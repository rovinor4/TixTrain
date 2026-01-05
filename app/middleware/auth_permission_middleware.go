package middleware

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// AuthPermissionMiddleware combines authentication and permission checking
func AuthPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. AUTH CHECK
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
		dataToken := pkg.DB.Where("value = ?", tokenString).First(&token)
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
		dataUser := pkg.DB.Where("id = ?", token.UserID).First(&user)
		if dataUser.RowsAffected == 0 || dataUser.Error != nil {
			c.JSON(401, gin.H{
				"message": "User tidak ditemukan",
			})
			c.Abort()
			return
		}

		// Get user roles from user_roles join table
		var userRoles []model.UserRole
		pkg.DB.Where("user_id = ?", user.ID).Find(&userRoles)

		// Extract role names
		var roleNames []string
		for _, ur := range userRoles {
			var role model.Role
			if err := pkg.DB.First(&role, ur.RoleID).Error; err == nil {
				roleNames = append(roleNames, role.Name)
			}
		}

		c.Set("user", user)
		c.Set("user_id", user.ID)
		c.Set("user_roles", roleNames)
		c.Set("token", tokenString)

		// 2. PERMISSION CHECK
		if len(roleNames) == 0 {
			c.JSON(403, gin.H{
				"message": "Akses ditolak - User tidak memiliki role",
			})
			c.Abort()
			return
		}

		// Get current request info
		method := c.Request.Method
		path := c.Request.URL.Path
		requestRoute := fmt.Sprintf("%s %s", method, path)

		// Fetch user's role permissions from database
		var rolePermissions []model.Permission

		err := pkg.DB.
			Joins("INNER JOIN role_permissions ON role_permissions.permission_id = permissions.id").
			Joins("INNER JOIN roles ON roles.id = role_permissions.role_id").
			Where("roles.name IN ?", roleNames).
			Find(&rolePermissions).Error

		if err != nil {
			c.JSON(403, gin.H{
				"message": "Akses ditolak",
			})
			c.Abort()
			return
		}

		// Check if current route matches any permission
		hasPermission := false
		for _, perm := range rolePermissions {
			if strings.TrimSpace(perm.Route) == strings.TrimSpace(requestRoute) {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(403, gin.H{
				"message":        "Akses ditolak - Anda tidak memiliki permission untuk route ini",
				"required_route": requestRoute,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

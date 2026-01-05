package middleware

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// PermissionMiddleware checks if user has permission for the specific route
func PermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoles, exists := c.Get("user_roles")
		if !exists {
			c.JSON(403, gin.H{
				"message": "Akses ditolak",
			})
			c.Abort()
			return
		}

		roles, ok := userRoles.([]string)
		if !ok || len(roles) == 0 {
			c.JSON(403, gin.H{
				"message": "Akses ditolak",
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
			Where("roles.name IN ?", roles).
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
			// Simple route matching: check if permission route matches current route
			// For example: "GET /admin/stats" matches "GET /admin/stats"
			if isRouteMatch(perm.Route, requestRoute) {
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

// isRouteMatch checks if a permission route matches the request route
func isRouteMatch(permissionRoute string, requestRoute string) bool {
	// Exact match
	if strings.TrimSpace(permissionRoute) == strings.TrimSpace(requestRoute) {
		return true
	}

	// For more complex matching with parameters, you can implement pattern matching here
	// For now, we do simple exact matching
	return false
}

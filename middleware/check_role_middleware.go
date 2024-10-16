package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CheckRoleMiddleware проверяет роль пользователя
func CheckRoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role") // Получение роли из контекста (предполагается, что она была установлена при аутентификации)

		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		roleString := strings.Join(roles, ",")
		if !contains(roleString, userRole.(string)) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// contains проверяет, содержится ли роль в списке
func contains(roleString string, role string) bool {
	roles := strings.Split(roleString, ",")
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

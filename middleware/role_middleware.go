package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleAuthMiddleware memverifikasi role pengguna
func RoleAuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil role dari context yang di-set oleh middleware
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get role"})
			return
		}

		// Periksa role
		if role == requiredRole {
			c.Next() // Lanjutkan ke handler berikutnya jika role sesuai
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission"})
			c.Abort()
		}
	}
}

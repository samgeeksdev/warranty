package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"warranty/api/authenticates"
)

func AdminOnly(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Extract user ID from Gin context (replace with your specific logic)
		userID, err := authenticates.ExtractUserID(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		isAdmin, err := authenticates.IsAdmin(db, userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if !isAdmin {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// Call the next handler in the chain if authorized
		c.Next()
	}
}

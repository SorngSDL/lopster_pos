package middlewares

import (
	"net/http"
	"strings"

	"com.lopster-pos/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "missing or invalid token"})
			return
		}

		token := strings.TrimSpace(auth[len("Bearer "):])

		userID, phone, err := utils.ValidateAccessToken(token)
		if err != nil || userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set("userId", userID)
		c.Set("phone", phone)

		c.Next()
	}
}

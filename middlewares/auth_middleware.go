package middlewares
import (
	"event-management/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		email, role, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("email", email)
		c.Set("role", role)
		c.Next()
	}
}

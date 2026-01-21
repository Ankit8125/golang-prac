package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
/abc -> only admin can access -> 2 level check -> auth -> 2nd (RequireAdmin)
/def -> any auth can access -> 1 level check -> auth
/ghi -> anyone can access -> no level check
*/

func RequireAdmin () gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := GetUserRole(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"error": "Unauthroized",
			})
			return 
		}

		if !strings.EqualFold(role, "admin"){
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H {
				"error": "This route can only be accessed by admin",
			})
			return 
		}

		c.Next()
	}
}
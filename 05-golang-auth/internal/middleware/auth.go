package middleware

import (
	"fmt"
	"go-auth/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Store -> auth data info -> gin context

const (
	ctxUserIDKey = "auth.userID"
	ctxRoleKey = "auth.role"
)

func AuthRequired (jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization token",
			})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2) // Authorization: Bearer <token>
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization format",
			})
			return
		}

		scheme := strings.TrimSpace(parts[0])
		tokenString := strings.TrimSpace(parts[1])

		if !strings.EqualFold(scheme, "Bearer"){
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization Scheme must be bearer",
			})
			return
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing token here",
			})
			return
		}

		claims, err := auth.ParseToken(jwtSecret, tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return 
		}

		c.Set(ctxUserIDKey, claims.Subject) // Setting this as key-value pair in gin.context
		c.Set(ctxRoleKey, claims.Role)

		c.Next()
	}
}

func GetUserID (c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxUserIDKey)
	if !ok {
		return "", false
	}

	userID, ok := res.(string)
	return userID, ok
}

func GetUserRole (c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxRoleKey)
	if !ok {
		return "", false
	}

	role, ok := res.(string)
	return role, ok
}
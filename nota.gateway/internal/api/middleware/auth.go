package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(protectedPrefixes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		isProtected := false
		path := c.Request.URL.Path

		for _, prefix := range protectedPrefixes {
			if strings.HasPrefix(path, prefix) {
				isProtected = true
				break
			}
		}

		if !isProtected {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		token := parts[1]
		ctx := context.WithValue(c.Request.Context(), "accessToken", token)

		userID := "dde009e4-aad0-4570-b40a-cb0caee2a1c1"
		ctx = context.WithValue(ctx, "userID", userID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

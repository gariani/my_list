package middleware

import (
	"net/http"

	"github.com/gariani/my_list/src/auth"
	"github.com/gin-gonic/gin"
)

func VerifyCSRF() gin.HandlerFunc {
	return func(c *gin.Context) {
		csrfCookie, err := c.Cookie("csrf_token")
		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{"error": "CSRF cookie missing"})
			return
		}

		csrfHeader := c.GetHeader("X-CSRF-Token")
		if csrfHeader == "" || csrfHeader != csrfCookie {
			c.AbortWithStatusJSON(403, gin.H{"error": "CSRF token invalid"})
			return
		}

		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}

		claims, valid := auth.ValidateAccessToken(token)
		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("userId", claims["user_id"])
		c.Next()
	}
}

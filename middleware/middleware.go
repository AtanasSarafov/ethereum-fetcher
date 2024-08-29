package middleware

import (
	"eth-fetcher/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware is a middleware that checks for a valid JWT token in the request headers.
func AuthMiddleware(authService services.AuthServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("AUTH_TOKEN")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token not provided"})
			c.Abort()
			return
		}

		// Validate the JWT token
		token, err := authService.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization token"})
			c.Abort()
			return
		}

		// Store the token claims in the context for further use
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
		}

		c.Next()
	}
}

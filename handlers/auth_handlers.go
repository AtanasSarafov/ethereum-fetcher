package handlers

import (
	"eth-fetcher/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthenticateHandler handles user authentication and JWT token generation.
func AuthenticateHandler(authService services.AuthServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// Bind JSON input to credentials struct
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		// Authenticate the user and get the JWT token
		token, err := authService.GenerateJWT(credentials.Username, credentials.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
			return
		}

		// Return the token as JSON response
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

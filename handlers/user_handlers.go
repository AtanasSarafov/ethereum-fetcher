package handlers

import (
	"eth-fetcher/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FetchUserTransactionsHandler handles fetching all transactions requested by the authenticated user.
func FetchUserTransactionsHandler(service *services.EthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// The username is set by the JWT middleware, so we can directly retrieve it
		username, _ := c.Get("username")

		authToken := c.GetHeader("AUTH_TOKEN")
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "AUTH_TOKEN is required"})
			return
		}

		transactions, err := service.GetUserTransactions(username.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user transactions"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	}
}

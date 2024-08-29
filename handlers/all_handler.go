package handlers

import (
	"eth-fetcher/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FetchAllTransactionsHandler handles fetching all transactions stored in the database.
func FetchAllTransactionsHandler(service *services.EthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch all transactions from the database
		transactions, err := service.GetAllTransactions()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transactions from database"})
			return
		}

		// Convert the database transaction models to the response format
		var response []gin.H
		for _, tx := range transactions {
			response = append(response, gin.H{
				"transactionHash":   tx.TransactionHash,
				"transactionStatus": tx.TransactionStatus,
				"blockHash":         tx.BlockHash,
				"blockNumber":       tx.BlockNumber,
				"from":              tx.From,
				"to":                tx.To,
				"contractAddress":   tx.ContractAddress,
				"logsCount":         tx.LogsCount,
				"input":             tx.Input,
				"value":             tx.Value,
			})
		}

		// Return the list of transactions as a JSON response
		c.JSON(http.StatusOK, gin.H{"transactions": response})
	}
}

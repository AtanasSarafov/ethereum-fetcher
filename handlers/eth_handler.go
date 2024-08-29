package handlers

import (
	"encoding/hex"
	"eth-fetcher/services"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/gin-gonic/gin"
)

// FetchTransactionsHandler handles fetching transactions based on a list of transaction hashes.
func FetchTransactionsHandler(service *services.EthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// The username is set by the JWT middleware, so we can directly retrieve it
		username, _ := c.Get("username")

		hashes := c.QueryArray("transactionHashes")
		if len(hashes) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "transactionHashes parameter is required"})
			return
		}

		var response []map[string]interface{}
		for _, hashStr := range hashes {
			txHash := common.HexToHash(hashStr)
			txDetails, err := service.ProcessTransaction(txHash)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transaction details"})
				return
			}
			response = append(response, txDetails)

			// Save the transaction hash for the user if authenticated
			if username != nil {
				service.SaveUserTransaction(username.(string), txHash.Hex())
			}
		}

		c.JSON(http.StatusOK, gin.H{"transactions": response})
	}
}

// FetchTransactionsByRLPHandler handles fetching transactions based on an RLP-encoded list of transaction hashes.
func FetchTransactionsByRLPHandler(service *services.EthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// The username is set by the JWT middleware, so we can directly retrieve it
		username, _ := c.Get("username")

		rlpHex := c.Param("rlphex")
		if rlpHex == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "rlphex parameter is required"})
			return
		}

		rlpBytes, err := hex.DecodeString(rlpHex)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid RLP hex string"})
			return
		}

		var txHashes []common.Hash
		if err := rlp.DecodeBytes(rlpBytes, &txHashes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to decode RLP data"})
			return
		}

		var response []map[string]interface{}
		for _, txHash := range txHashes {
			txDetails, err := service.ProcessTransaction(txHash)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transaction details"})
				return
			}
			response = append(response, txDetails)

			// Save the transaction hash for the user if authenticated
			if username != nil {
				service.SaveUserTransaction(username.(string), txHash.Hex())
			}
		}

		c.JSON(http.StatusOK, gin.H{"transactions": response})
	}
}

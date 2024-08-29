package eth

// import (
// 	"context"
// 	"eth-fetcher/models"
// 	"log"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"gorm.io/gorm"
// )

// func FetchTransactions(db *gorm.DB, transactionHashes []string) []models.Transaction {
// 	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var transactions []models.Transaction
// 	for _, hash := range transactionHashes {
// 		var tx models.Transaction
// 		db.First(&tx, "transaction_hash = ?", hash)
// 		if tx.TransactionHash != "" {
// 			transactions = append(transactions, tx)
// 			continue
// 		}

// 		txHash := common.HexToHash(hash)
// 		txData, _, err := client.TransactionByHash(context.Background(), txHash)
// 		if err != nil {
// 			log.Println("Transaction not found:", err)
// 			continue
// 		}

// 		receipt, err := client.TransactionReceipt(context.Background(), txHash)
// 		if err != nil {
// 			log.Println("Receipt not found:", err)
// 			continue
// 		}

// 		tx = models.Transaction{
// 			TransactionHash:   txData.Hash().Hex(),
// 			TransactionStatus: receipt.Status,
// 			BlockHash:         receipt.BlockHash.Hex(),
// 			BlockNumber:       receipt.BlockNumber.Uint64(),
// 			From:              txData.From().Hex(),
// 			To:                txData.To().Hex(),
// 			ContractAddress:   receipt.ContractAddress.Hex(),
// 			LogsCount:         len(receipt.Logs),
// 			Input:             string(txData.Data()),
// 			Value:             txData.Value().String(),
// 		}

// 		db.Create(&tx)
// 		transactions = append(transactions, tx)
// 	}
// 	return transactions
// }

package services

import (
	"context"
	"eth-fetcher/models"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// EthServiceInterface defines the methods for interacting with Ethereum transactions.
type EthServiceInterface interface {
	GetAllTransactions() ([]models.Transaction, error)
	FetchTransactionByHash(txHash string) (*models.Transaction, error)
	SaveTransaction(tx *models.Transaction) error
	RecordUserTransaction(userID string, txHash string) error
	GetTransactionsByUser(userID string) ([]models.Transaction, error)
}

// EthService -
type EthService struct {
	client *ethclient.Client
	db     *gorm.DB
}

// NewEthService -
func NewEthService(ethereumNodeURL string, db *gorm.DB) (*EthService, error) {
	client, err := ethclient.Dial(ethereumNodeURL)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}

	return &EthService{
		client: client,
		db:     db,
	}, nil
}

// GetTransactionDetails fetches the transaction, receipt, and chain ID.
func (e *EthService) GetTransactionDetails(txHash common.Hash) (*types.Transaction, *types.Receipt, *big.Int, error) {
	fmt.Println("!!! GetTransactionDetails: txHash: " + txHash.String())

	tx, isPending, err := e.client.TransactionByHash(context.Background(), txHash)
	fmt.Println("!!! GetTransactionDetails: 1 err:!!!" + err.Error())

	if err != nil || isPending {
		return nil, nil, nil, err
	}

	fmt.Println("!!! GetTransactionDetails: 2 !!!")

	receipt, err := e.client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return nil, nil, nil, err
	}

	chainID, err := e.client.NetworkID(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}

	return tx, receipt, chainID, nil
}

// GetBlockDetails fetches block details by block hash.
func (e *EthService) GetBlockDetails(blockHash common.Hash) (*types.Block, error) {
	block, err := e.client.BlockByHash(context.Background(), blockHash)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// SaveUserTransaction saves the transaction hash for a specific user.
func (e *EthService) SaveUserTransaction(username string, txHash string) error {
	userTx := models.UserTransaction{
		Username: username,
		TxHash:   txHash,
	}

	return e.db.Create(&userTx).Error
}

// GetUserTransactions retrieves all transactions for a specific user.
func (e *EthService) GetUserTransactions(username string) ([]map[string]interface{}, error) {
	var userTxs []models.UserTransaction
	err := e.db.Where("username = ?", username).Find(&userTxs).Error
	if err != nil {
		return nil, err
	}

	var transactions []map[string]interface{}
	for _, userTx := range userTxs {
		txHash := common.HexToHash(userTx.TxHash)
		txDetails, err := e.ProcessTransaction(txHash)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, txDetails)
	}

	return transactions, nil
}

// ProcessTransaction handles fetching transaction details, recovering the sender address, and assembling the response.
func (e *EthService) ProcessTransaction(txHash common.Hash) (map[string]interface{}, error) {
	tx, receipt, chainID, err := e.GetTransactionDetails(txHash)
	fmt.Println("!!! 1 !!!")

	if err != nil {
		return nil, err
	}
	fmt.Println("!!! 2 !!!")

	// Recover the sender's address
	signer := types.NewEIP155Signer(chainID)
	from, err := types.Sender(signer, tx)
	if err != nil {
		return nil, err
	}

	fmt.Println("!!! 3 !!!")

	// Create the response map
	response := map[string]interface{}{
		"transactionHash":   tx.Hash().Hex(),
		"transactionStatus": receipt.Status,
		"blockHash":         receipt.BlockHash.Hex(),
		"blockNumber":       receipt.BlockNumber.Uint64(),
		"from":              from.Hex(),
		"to":                tx.To().Hex(),
		"contractAddress":   receipt.ContractAddress.Hex(),
		"logsCount":         len(receipt.Logs),
		"input":             tx.Data(),
		"value":             tx.Value().String(),
	}

	return response, nil
}

// GetAllTransactions fetches all transactions stored in the database.
// func (e *EthService) GetAllTransactions() ([]models.Transaction, error) {
// 	// Prepare SQL query to fetch all transactions
// 	rows, err := e.db.Query(`SELECT transaction_hash, transaction_status, block_hash, block_number, "from", "to", contract_address, logs_count, input, value FROM transactions`)
// 	if err != nil {
// 		log.Printf("Failed to query transactions: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	// Iterate through the result set and populate the transactions slice
// 	var transactions []models.Transaction
// 	for rows.Next() {
// 		var tx models.Transaction
// 		if err := rows.Scan(&tx.TransactionHash, &tx.TransactionStatus, &tx.BlockHash, &tx.BlockNumber, &tx.From, &tx.To, &tx.ContractAddress, &tx.LogsCount, &tx.Input, &tx.Value); err != nil {
// 			log.Printf("Failed to scan transaction row: %v", err)
// 			return nil, err
// 		}
// 		transactions = append(transactions, tx)
// 	}

// 	// Check for any errors encountered during iteration
// 	if err := rows.Err(); err != nil {
// 		log.Printf("Row iteration error: %v", err)
// 		return nil, err
// 	}

// 	return transactions, nil
// }

// GetAllTransactions fetches all transactions stored in the database using GORM.
func (e *EthService) GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction

	// Use GORM to find all transactions in the database
	err := e.db.Find(&transactions).Error
	if err != nil {
		log.Printf("Failed to fetch transactions: %v", err)
		return nil, err
	}

	return transactions, nil
}

// Close shuts down the Ethereum client connection.
func (e *EthService) Close() {
	e.client.Close()
}

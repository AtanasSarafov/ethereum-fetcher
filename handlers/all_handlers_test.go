package handlers

// import (
// 	"eth-fetcher/generated/mocks"
// 	"eth-fetcher/models"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestFetchAllTransactionsHandler(t *testing.T) {
// 	// Create a new mock service using the generated mock
// 	mockService := new(mocks.MockEthServiceInterface)

// 	// Define the mock transactions that should be returned by the service
// 	mockTransactions := []models.Transaction{
// 		{
// 			TransactionHash:   "0x123",
// 			TransactionStatus: 1,
// 			BlockHash:         "0xabc",
// 			BlockNumber:       123456,
// 			From:              "0xSenderAddress",
// 			To:                "0xReceiverAddress",
// 			ContractAddress:   "",
// 			LogsCount:         2,
// 			Input:             "0xInputData",
// 			Value:             "1000000000000000000", // 1 ETH in wei
// 		},
// 		{
// 			TransactionHash:   "0x456",
// 			TransactionStatus: 0,
// 			BlockHash:         "0xdef",
// 			BlockNumber:       123457,
// 			From:              "0xAnotherSender",
// 			To:                "0xAnotherReceiver",
// 			ContractAddress:   "0xNewContractAddress",
// 			LogsCount:         5,
// 			Input:             "0xMoreInputData",
// 			Value:             "500000000000000000", // 0.5 ETH in wei
// 		},
// 	}

// 	// Setup the mock to return the mock transactions
// 	mockService.On("GetAllTransactions").Return(mockTransactions, nil)

// 	// Initialize the Gin router and the handler
// 	router := gin.Default()
// 	router.GET("/lime/all", FetchAllTransactionsHandler(mockService))

// 	// Create a new HTTP request to test the handler
// 	req, _ := http.NewRequest(http.MethodGet, "/lime/all", nil)
// 	w := httptest.NewRecorder()

// 	// Serve the request
// 	router.ServeHTTP(w, req)

// 	// Assert the response status code
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// Expected JSON response
// 	expectedResponse := `{
// 		"transactions": [
// 			{
// 				"transactionHash": "0x123",
// 				"transactionStatus": 1,
// 				"blockHash": "0xabc",
// 				"blockNumber": 123456,
// 				"from": "0xSenderAddress",
// 				"to": "0xReceiverAddress",
// 				"contractAddress": "",
// 				"logsCount": 2,
// 				"input": "0xInputData",
// 				"value": "1000000000000000000"
// 			},
// 			{
// 				"transactionHash": "0x456",
// 				"transactionStatus": 0,
// 				"blockHash": "0xdef",
// 				"blockNumber": 123457,
// 				"from": "0xAnotherSender",
// 				"to": "0xAnotherReceiver",
// 				"contractAddress": "0xNewContractAddress",
// 				"logsCount": 5,
// 				"input": "0xMoreInputData",
// 				"value": "500000000000000000"
// 			}
// 		]
// 	}`

// 	// Assert the JSON response matches the expected response
// 	assert.JSONEq(t, expectedResponse, w.Body.String())

// 	// Assert that the mock service's GetAllTransactions method was called once
// 	mockService.AssertExpectations(t)
// }

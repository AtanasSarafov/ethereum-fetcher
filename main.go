package main

import (
	"eth-fetcher/config"
	"eth-fetcher/database"
	"eth-fetcher/handlers"
	"eth-fetcher/middleware"
	"eth-fetcher/services"

	"github.com/gin-gonic/gin"
)

// main is the entry point of the application. It initializes the configuration, database, and services,
// sets up the API routes, and starts the HTTP server.
func main() {
	// Load application configuration
	cfg := config.LoadConfig()

	// Initialize the database connection
	db := database.InitDB(cfg)

	// Create an Ethereum service instance, which handles interaction with Ethereum and the database
	ethService, err := services.NewEthService(cfg.EthereumNodeURL, db)
	if err != nil {
		// TODO: Handle the error.
	}

	authService := services.NewAuthService(cfg.JWTSecret)

	// Initialize Gin router
	r := gin.Default()

	// Authentication route
	r.POST("/lime/authenticate", handlers.AuthenticateHandler(authService))

	// Define the API routes and assign the corresponding handler functions
	// Fetch transactions routes
	r.GET("/lime/eth", middleware.AuthMiddleware(authService), handlers.FetchTransactionsHandler(ethService))
	r.GET("/lime/eth/:rlphex", middleware.AuthMiddleware(authService), handlers.FetchTransactionsByRLPHandler(ethService))
	r.GET("/lime/all", handlers.FetchAllTransactionsHandler(ethService))

	// User-specific transactions route
	r.GET("/lime/my", middleware.AuthMiddleware(authService), handlers.FetchUserTransactionsHandler(ethService))

	// Start the server on the specified port (default is 8080)
	r.Run(":" + cfg.APIPort)
}

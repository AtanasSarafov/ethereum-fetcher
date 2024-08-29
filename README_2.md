export API_PORT=8080
export ETH_NODE_URL=https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID
export DB_CONNECTION_URL=postgres://user:password@localhost:5432/yourdbname?sslmode=disable
export JWT_SECRET=your_jwt_secret

go run .


docker run -p 8080:8080 -e API_PORT=8080 -e ETH_NODE_URL="http://your-eth-node-url" -e DB_CONNECTION_URL="your-db-url" -e JWT_SECRET="your-secret-key" limeapi

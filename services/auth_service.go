package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

//var jwtSecret = []byte("your_jwt_secret_key")

// Claims represents the JWT claims.
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthServiceInterface interface {
	GenerateJWT(username, password string) (string, error)
	ValidateJWT(tokenString string) (*jwt.Token, error)
}

type AuthService struct {
	jwtSecret string
}

func NewAuthService(jwtSecret string) AuthServiceInterface {
	return &AuthService{
		jwtSecret: jwtSecret,
	}
}

// GenerateJWT generates a JWT token for a given username.
func (s *AuthService) GenerateJWT(username, password string) (string, error) {
	// Validate the username and password
	if (username == "alice" && password == "alice") || (username == "bob" && password == "bob") || (username == "carol" && password == "carol") || (username == "dave" && password == "dave") {
		// Generate the JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		})

		// Sign the token with the secret
		tokenString, err := token.SignedString([]byte(s.jwtSecret))
		if err != nil {
			return "", err
		}

		return tokenString, nil
	}

	return "", errors.New("invalid username or password")
}

// ValidateJWT validates a JWT token string and returns the parsed token if valid.
func (s *AuthService) ValidateJWT(tokenString string) (*jwt.Token, error) {
	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})

	// Return error if token is invalid
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, nil
	}

	return nil, errors.New("invalid token")
}

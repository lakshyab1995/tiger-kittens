package jwt

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenModel struct {
	Token  string
	Expiry string
}

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(username string) (*TokenModel, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	expiry := time.Now().Add(time.Hour * 24).Unix()
	claims["username"] = username
	claims["exp"] = expiry
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return nil, err
	}
	return &TokenModel{
		Token:  tokenString,
		Expiry: fmt.Sprintf("%d", expiry),
	}, nil
}

// ParseToken parses a jwt token and returns the username in it's claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		secretKey := []byte(os.Getenv("SECRET_KEY"))
		return secretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}

func IsTokenExpired(tokenStr string) bool {
	_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		secretKey := []byte(os.Getenv("SECRET_KEY"))
		return secretKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return true
			}
		}
	}
	return false
}

package auth

import (
	"encoding/base64"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	Id       string
	Name     string
	Username string
	jwt.RegisteredClaims
}

func (claim *UserClaim) SignAuthToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, _ := token.SignedString(secretKey)
	return base64.StdEncoding.EncodeToString([]byte(tokenString)), nil
}

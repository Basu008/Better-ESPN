package auth

import (
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
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, _ := token.SignedString(secretKey)
	return tokenString, nil
}

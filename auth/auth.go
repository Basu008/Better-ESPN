package auth

import (
	"encoding/base64"
	"fmt"
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
	fmt.Println("Token : ", base64.StdEncoding.EncodeToString([]byte(tokenString)))
	d, _ := VerifyAuthToken(base64.StdEncoding.EncodeToString([]byte(tokenString)))
	fmt.Println("Back To Normal :", d)
	return base64.StdEncoding.EncodeToString([]byte(tokenString)), nil
}

func VerifyAuthToken(tokenString string) (UserClaim, error) {
	claim := UserClaim{}
	decodedString, err := base64.StdEncoding.DecodeString(tokenString)
	if err != nil {
		return claim, err
	}
	token, err := jwt.ParseWithClaims(string(decodedString), &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return claim, err
	}
	if !token.Valid {
		return claim, err
	}
	return claim, err
}

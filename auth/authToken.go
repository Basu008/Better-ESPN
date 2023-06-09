package auth

import (
	"encoding/base64"

	"github.com/Basu008/Better-ESPN/config"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (claim *UserClaim) SignUserToken(c *config.Config) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	tokenString, _ := token.SignedString([]byte(c.GetJWTKey()))
	return base64.StdEncoding.EncodeToString([]byte(tokenString)), nil
}

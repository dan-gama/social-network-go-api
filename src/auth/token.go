package auth

import (
	"sn-api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken cria um token com as permissões do usuário
func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 1).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.TokenSecret)
}

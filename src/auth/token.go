package auth

import (
	"errors"
	"fmt"
	"net/http"
	"sn-api/src/config"
	"strconv"
	"strings"
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

// ValidateToken verifica se o token na requisição é válido
func ValidateToken(r *http.Request) error {
	requestToken := getToken(r)
	token, err := jwt.Parse(requestToken, getSecretKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return errors.New("token inválido")
	}

	return nil
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func getSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("invalid signing method")
	}
	return config.TokenSecret, nil
}

// GetUserId retorna o usuário Id
func GetUserId(r *http.Request) (uint64, error) {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, getSecretKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userId, nil
	}

	return 0, errors.New("token inválido")
}

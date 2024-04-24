package middlewares

import (
	"fmt"
	"net/http"
	"sn-api/src/auth"
	"sn-api/src/responses"
)

// Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(
			"\n %s %s %s",
			r.Method,
			r.RequestURI,
			r.Host,
		)
		next(w, r)
	}
}

// Authenticate é o middleware responsável por validar o token do usuário
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

package security

import "golang.org/x/crypto/bcrypt"

// Hash aplica um hash na string recebida
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPassword verifica se a senha está correta
func CheckPassword(passwordWithHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(password))
}

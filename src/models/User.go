package models

import (
	"errors"
	"strings"
	"time"
)

// User representa a tabela User
type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Validate valida se a model está correta
func (user *User) Validate() error {
	if err := user.rules(); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) rules() error {
	if user.Name == "" {
		return errors.New("nome é obrigatório")
	}

	if user.Email == "" {
		return errors.New("e-mail é obrigatório")
	}

	if user.Password == "" {
		return errors.New("senha é obrigatório")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Name)
}

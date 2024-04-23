package models

import (
	"errors"
	"sn-api/src/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User representa a tabela User
type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

const (
	ActionCreate = iota
	ActionUpdate
)

// Validate valida se a model está correta
func (user *User) Validate(action int) error {
	if err := user.rules(action); err != nil {
		return err
	}

	if err := user.format(action); err != nil {
		return err
	}
	return nil
}

func (user *User) rules(action int) error {
	if user.Name == "" {
		return errors.New("nome é obrigatório")
	}

	if user.Email == "" {
		return errors.New("e-mail é obrigatório")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("e-mail com formtação inválida")
	}

	if action == ActionCreate && user.Password == "" {
		return errors.New("senha é obrigatório")
	}

	return nil
}

func (user *User) format(action int) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if action == ActionCreate {
		passwordWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordWithHash)
	}

	return nil
}

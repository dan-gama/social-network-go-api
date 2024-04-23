package repositories

import (
	"database/sql"
	"sn-api/src/models"
)

type login struct {
	db *sql.DB
}

// NewUserRepository cria um repositório de User
func NewLoginRepository(db *sql.DB) *login {
	return &login{db}
}

// Create cria um usuário no banco de dados
func (repo login) CheckAuth(email string) (models.User, error) {
	rows, err := repo.db.Query(`
		select
			id,
			password
		from
			User
		where
			email = ?
	`, email)

	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

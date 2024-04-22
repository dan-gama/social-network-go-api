package repositories

import (
	"database/sql"
	"sn-api/src/models"
)

type user struct {
	db *sql.DB
}

// NewUserRepository cria um repositório de User
func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

// Create cria um usuário no banco de dados
func (repo user) Create(user models.User) (uint64, error) {
	// return 0, nil
	statement, err := repo.db.Prepare(`
		insert into User (
			name,
			email, 
			password
		) values (?,?,?)
	`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	ret, err := statement.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

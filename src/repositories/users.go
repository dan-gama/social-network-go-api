package repositories

import (
	"database/sql"
	"fmt"
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
	statement, err := repo.db.Prepare(`
		insert into User (
			email,
			name, 
			password
		) values (?,?,?)
	`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	ret, err := statement.Exec(user.Email, user.Name, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// GetUsersByName busca todos os usuários
func (repo user) GetAll() ([]models.User, error) {
	rows, err := repo.db.Query(`
		select
			id,
			name,
			email,
			createdAt
		from
			User
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUsersByName busca os usuários pelo nome
func (repo user) GetByName(name string) ([]models.User, error) {
	rows, err := repo.db.Query(`
		select
			id,
			name,
			email,
			createdAt
		from
			User
		where
			name LIKE ? 
	`, fmt.Sprintf("%%%s%%", name))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Get retorna um usuário, de acordo com o id
func (repo user) Get(id uint64) (models.User, error) {
	rows, err := repo.db.Query(`
		select
			id,
			name,
			email,
			createdAt
		from
			User
		where
			id = ?
	`, id)

	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update Atualiza o usuário no banco de dados
func (repo user) Update(id uint64, model models.User) error {
	statement, err := repo.db.Prepare(`
		update User set
			email = ?,
			name = ?
		where
			id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(model.Email, model.Name, id); err != nil {
		return err
	}

	return nil
}

// Delete Exclui o usuário do banco de dados
func (repo user) Delete(id uint64) error {
	statement, err := repo.db.Prepare(`
		delete from User
		where
			id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

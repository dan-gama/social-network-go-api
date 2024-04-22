package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"sn-api/src/data"
	"sn-api/src/models"
	"sn-api/src/repositories"
	"sn-api/src/responses"
)

// Create cria um novo usuário
func Create(writer http.ResponseWriter, req *http.Request) {
	// write.Write([]byte("Creating user..."))
	body, err := io.ReadAll(req.Body)
	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	if err := user.Validate(); err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	db, err := data.ConnectDB()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.Id, err = repository.Create(user)
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(writer, http.StatusCreated, user)
}

// GetAll busca todos os usuários
func GetAll(write http.ResponseWriter, req *http.Request) {
	write.Write([]byte("Get all users..."))
}

// Get busca um determinado usuário
func Get(write http.ResponseWriter, req *http.Request) {
	write.Write([]byte("Get user..."))
}

// Update atualiza um usuário
func Update(write http.ResponseWriter, req *http.Request) {
	write.Write([]byte("Updating user..."))
}

// Delete remove um usuário
func Delete(write http.ResponseWriter, req *http.Request) {
	write.Write([]byte("Deleting user..."))
}

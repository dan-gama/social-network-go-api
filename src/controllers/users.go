package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sn-api/src/data"
	"sn-api/src/models"
)

// Create cria um novo usuário
func Create(write http.ResponseWriter, req *http.Request) {
	// write.Write([]byte("Creating user..."))
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := data.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
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

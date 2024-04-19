package controllers

import "net/http"

// Create cria um novo usuário
func Create(write http.ResponseWriter, req *http.Request) {
	write.Write([]byte("Creating user..."))
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

package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sn-api/src/data"
	"sn-api/src/models"
	"sn-api/src/repositories"
	"sn-api/src/responses"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type ModelValidate interface {
	Validate() error
}

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

	if err := user.Validate(models.ActionCreate); err != nil {
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
func GetAll(writer http.ResponseWriter, req *http.Request) {
	db, err := data.ConnectDB()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, err := repository.GetAll()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
	}

	responses.JSON(writer, http.StatusOK, users)
	// write.Write([]byte("Get all users..."))
}

// GetAll busca todos os usuários pelo nome
func GetByName(writer http.ResponseWriter, req *http.Request) {
	name := strings.ToLower(req.URL.Query().Get("name"))

	db, err := data.ConnectDB()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, err := repository.GetByName(name)
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
	}

	responses.JSON(writer, http.StatusOK, users)
	// write.Write([]byte("Get all users..."))
}

// Get busca um determinado usuário
func Get(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	db, err := data.ConnectDB()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	user, err := repository.Get(id)
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	if user.Id == 0 {
		responses.Error(writer, http.StatusNotFound, errors.New("usuário não encontrado"))
		return
	}

	responses.JSON(writer, http.StatusOK, user)
}

// Update atualiza um usuário
func Update(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
	}

	if err := user.Validate(models.ActionCreate); err != nil {
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
	err = repository.Update(id, user)
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(writer, http.StatusNoContent, nil)

	//writer.Write([]byte("Updating user..."))
}

// Delete remove um usuário
func Delete(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
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
	err = repository.Delete(id)
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(writer, http.StatusNoContent, nil)
}

func BodyToModel(req *http.Request, model ModelValidate) responses.ResponseError {
	var responseError responses.ResponseError

	// Le o corpo da requisição
	body, err := io.ReadAll(req.Body)
	if err != nil {
		responseError.StatusCode = http.StatusUnprocessableEntity
		responseError.Error = err
		return responseError
	}

	// Transforma o corpo da requisição num json, de acordo com a model
	if err = json.Unmarshal(body, model); err != nil {
		responseError.StatusCode = http.StatusBadRequest
		responseError.Error = err
		return responseError
	}

	// Valida a model
	if err := model.Validate(); err != nil {
		responseError.StatusCode = http.StatusBadRequest
		responseError.Error = err
		return responseError
	}

	return responses.ResponseError{}
}

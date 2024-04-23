package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sn-api/src/auth"
	"sn-api/src/data"
	"sn-api/src/models"
	"sn-api/src/repositories"
	"sn-api/src/responses"
	"sn-api/src/security"
)

func LoginAuth(writer http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		responses.Error(writer, http.StatusBadRequest, err)
		return
	}

	db, err := data.ConnectDB()
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewLoginRepository(db)
	userCheck, err := repository.CheckAuth(user.Email)
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(userCheck.Password, user.Password); err != nil {
		responses.Error(writer, http.StatusUnauthorized, errors.New("a senhha informada está incorreta"))
		return
	}

	token, err := auth.CreateToken(userCheck.Id)
	if err != nil {
		responses.Error(writer, http.StatusInternalServerError, errors.New("fallha ao gerar token"))
		return
	}

	writer.Write([]byte(token))

	// responses.JSON(writer, http.StatusOK, auth)
}

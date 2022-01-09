package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"encoding/json"
	"net/http"
)

type userHandlerImpl struct {
	AccountRepository model.AccountRepository
	UserService       port.UserService
}

func NewUserhandler(repo model.AccountRepository, userService port.UserService) port.UserHandler {
	return userHandlerImpl{AccountRepository: repo, UserService: userService}
}

func (handler userHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := dto.UserWithPasswordDto{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithJSON(w, 500, "error")
		return
	}
	basicUser := handler.UserService.FindBasicUser(user.Username)
	if basicUser != (dto.BasicUserDto{}) {
		respondWithJSON(w, 400, "username already exists")
		return
	}
	//TODO VALIDATE IF ACCOUNT WITH ID DOCUMENT ALREADY EXISTS
	//TODO MODIF ACCOUNT SERVICE FOR CREATE ACCOUNT
	_, err = handler.AccountRepository.CreateAccount(model.Account{AccountNumber: user.IDDocument})

	if err != nil {
		respondWithJSON(w, 500, "error")
		return
	}
	respondWithJSON(w, 200, "ok")
}

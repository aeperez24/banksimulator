package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"encoding/json"
	"net/http"
	"reflect"
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
	//TODO MODIF ACCOUNT SERVICE FOR CREATE ACCOUNT

	acc := handler.AccountRepository.FindAccountByAccountNumber(user.IDDocument)

	if !reflect.DeepEqual(acc, model.Account{}) {
		respondWithJSON(w, 400, "account already exists")
		return
	}

	_, err = handler.AccountRepository.CreateAccount(model.Account{AccountNumber: user.IDDocument})

	if err != nil {
		respondWithJSON(w, 500, "error creating account")
		return
	}
	err = handler.UserService.CreateUser(user)

	if err != nil {
		respondWithJSON(w, 500, "error creating user")
		return
	}

	respondWithJSON(w, 200, "ok")
}

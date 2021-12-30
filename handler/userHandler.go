package handler

import (
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"net/http"

	"github.com/gorilla/mux"
)

type userHandlerImpl struct {
	AccountRepository model.AccountRepository
	UserService       port.UserService
}

func NewUserhandler(repo model.AccountRepository, userService port.UserService) port.UserHandler {
	return userHandlerImpl{AccountRepository: repo, UserService: userService}
}

func (handler userHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	respondWithJSON(w, 200, vars)
}

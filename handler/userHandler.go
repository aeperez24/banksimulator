package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/port"
	"aeperez24/banksimulator/usercase"
	"encoding/json"
	"net/http"
)

type userHandlerImpl struct {
	usercase.UserUsercase
}

func NewUserhandler(useCase usercase.UserUsercase) port.UserHandler {
	return userHandlerImpl{useCase}
}

func (handler userHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := dto.UserWithPasswordDto{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithJSON(w, 500, "error")
		return
	}
	usecaseError := handler.UserUsercase.CreateUser(user)
	if usecaseError != (usercase.UserCaseError{}) {
		respondWithJSON(w, usecaseError.Code, usecaseError.Message)
		return
	}

	respondWithJSON(w, 200, "ok")
}

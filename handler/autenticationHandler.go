package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/port"
	"encoding/json"
	"errors"
	"net/http"
)

type autenticationHandlerImpl struct {
	userService port.UserService
	tkservice   port.TokenService
}

func (handler autenticationHandlerImpl) Authenticate(w http.ResponseWriter, r *http.Request) {
	userDto := dto.UserWithPasswordDto{}
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		respondWithJSON(w, 400, err)
		return
	}

	token, err := handler.ExecuteAuthenticaton(userDto)
	if err != nil {
		respondWithJSON(w, 400, err)
		return
	}
	respondWithJSON(w, 200, token)
}

func (handler autenticationHandlerImpl) ExecuteAuthenticaton(userdto dto.UserWithPasswordDto) (string, error) {
	valid := handler.userService.ValidateUserameAndPassword(userdto.Username, userdto.Password)
	if !valid {
		return "", errors.New("invalid username or password")
	}

	user := handler.userService.FindBasicUser(userdto.Username)
	if user == (dto.BasicUserDto{}) {
		return "", errors.New("user not found")
	}

	return handler.tkservice.CreateToken(user)
}

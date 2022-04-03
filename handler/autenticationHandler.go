package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/port"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

type GinAutenticationHandlerImpl struct {
	userService port.UserService
	tkservice   port.TokenService
}

func NewAuthenticationHandlerGin(userService port.UserService, tkservice port.TokenService) GinAutenticationHandlerImpl {
	return GinAutenticationHandlerImpl{userService, tkservice}
}

func (handler GinAutenticationHandlerImpl) Authenticate(c *gin.Context) {
	userDto := dto.UserWithPasswordDto{}
	err := c.ShouldBindJSON(&userDto)
	log.Printf("user %v", userDto)
	if err != nil {
		log.Fatal(err)
		respondWithJSONGin(c, 400, "")
		return
	}

	token, err := handler.ExecuteAuthenticaton(userDto)
	log.Printf("token %v", err)
	if err != nil {
		log.Fatal(err)
		respondWithJSONGin(c, 400, "")
		return
	}
	respondWithJSONGin(c, 200, token)
}

func (handler GinAutenticationHandlerImpl) ExecuteAuthenticaton(userdto dto.UserWithPasswordDto) (string, error) {
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

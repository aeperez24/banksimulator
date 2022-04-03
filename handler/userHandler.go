package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/usercase"

	"github.com/gin-gonic/gin"
)

type GinUserHandlerImpl struct {
	usercase.UserUsercase
}

func NewGinUserhandler(useCase usercase.UserUsercase) GinUserHandlerImpl {
	return GinUserHandlerImpl{useCase}
}

func (handler GinUserHandlerImpl) CreateUser(c *gin.Context) {
	user := dto.UserWithPasswordDto{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		respondWithJSONGin(c, 500, "error")
		return
	}
	usecaseError := handler.UserUsercase.CreateUser(user)
	if usecaseError != (usercase.UserCaseError{}) {
		respondWithJSONGin(c, usecaseError.Code, usecaseError.Message)
		return
	}

	respondWithJSONGin(c, 200, "ok")
}

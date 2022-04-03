package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/port"
	"aeperez24/banksimulator/usercase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinAccountHandlerImpl struct {
	AccountUsercase usercase.AccountUsercase
}

func NewGinAccountHandler(AccountUsercase usercase.AccountUsercase) GinAccountHandlerImpl {
	return GinAccountHandlerImpl{AccountUsercase}
}

func (handler GinAccountHandlerImpl) GetBalance(c *gin.Context) {
	log.Printf("user pt3 :%v", c.Request.Context().Value(port.LoggedUserKey))

	balance, err := handler.AccountUsercase.GetBalance(c.Param("accountNumber"))
	if err != (usercase.UserCaseError{}) {
		c.JSON(err.Code, gin.H{"data": err.Message})
		respondWithJSONGin(c, err.Code, err.Message)
	}
	respondWithJSONGin(c, 200, balance)
}

func (handler GinAccountHandlerImpl) GetTransactions(c *gin.Context) {
	ctx := c.Request.Context()
	transactions, err := handler.AccountUsercase.GetTransactions(ctx, c.Param("accountNumber"))
	if err != (usercase.UserCaseError{}) {
		respondWithJSONGin(c, err.Code, err.Message)
	}
	respondWithJSONGin(c, 200, transactions)

}

func (handler GinAccountHandlerImpl) TransferMoney(c *gin.Context) {
	var transferRequest dto.TransferRequest
	err := c.ShouldBindJSON(&transferRequest)
	if err != nil {
		log.Fatal(err)
		respondWithJSONGin(c, http.StatusBadRequest, "")
		return
	}
	balance, usecaseError := handler.AccountUsercase.TransferMoney(c.Request.Context(), transferRequest)
	if usecaseError != (usercase.UserCaseError{}) {
		respondWithJSONGin(c, usecaseError.Code, usecaseError.Message)
	}
	respondWithJSONGin(c, 200, balance)

}

func (handler GinAccountHandlerImpl) Deposit(c *gin.Context) {
	var depositRequest dto.DepositRequest
	err := c.ShouldBindJSON(&depositRequest)
	if err != nil {
		log.Fatal(err)
		respondWithJSONGin(c, http.StatusBadRequest, "")
		return
	}
	balance, usecaseError := handler.AccountUsercase.Deposit(c.Request.Context(), depositRequest)
	if usecaseError != (usercase.UserCaseError{}) {
		respondWithJSONGin(c, usecaseError.Code, usecaseError.Message)
	}
	respondWithJSONGin(c, 200, balance)
}

package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/port"
	"aeperez24/banksimulator/usercase"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type accountHandlerImpl struct {
	AccountUsercase usercase.AccountUsercase
}

func NewAccountHandler(AccountUsercase usercase.AccountUsercase) port.AccountHandler {
	return accountHandlerImpl{AccountUsercase}
}

func (handler accountHandlerImpl) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	balance, err := handler.AccountUsercase.GetBalance(vars["AccountNumber"])
	if err != (usercase.UserCaseError{}) {
		respondWithJSON(w, err.Code, err.Message)
	}
	respondWithJSON(w, 200, balance)
}

func (handler accountHandlerImpl) GetTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactions, err := handler.AccountUsercase.GetTransactions(r.Context(), vars["AccountNumber"])
	if err != (usercase.UserCaseError{}) {
		respondWithJSON(w, err.Code, err.Message)
	}
	respondWithJSON(w, 200, transactions)

}

func (handler accountHandlerImpl) TransferMoney(w http.ResponseWriter, r *http.Request) {
	var transferRequest dto.TransferRequest
	err := json.NewDecoder(r.Body).Decode(&transferRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	balance, usecaseError := handler.AccountUsercase.TransferMoney(r.Context(), transferRequest)
	if usecaseError != (usercase.UserCaseError{}) {
		respondWithJSON(w, usecaseError.Code, usecaseError.Message)
	}
	respondWithJSON(w, 200, balance)

}

func (handler accountHandlerImpl) Deposit(w http.ResponseWriter, r *http.Request) {
	var depositRequest dto.DepositRequest
	err := json.NewDecoder(r.Body).Decode(&depositRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	balance, usecaseError := handler.AccountUsercase.Deposit(r.Context(), depositRequest)
	if usecaseError != (usercase.UserCaseError{}) {
		respondWithJSON(w, usecaseError.Code, usecaseError.Message)
	}
	respondWithJSON(w, 200, balance)
}

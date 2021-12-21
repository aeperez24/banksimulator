package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"aeperez24/banksimulator/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type accountHandlerImpl struct {
	AccountRepository model.AccountRepository
}

func NewAccountHandler(repo model.AccountRepository) port.AccountHandler {
	return accountHandlerImpl{repo}
}

func (handler accountHandlerImpl) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := handler.getAccountService(vars["AccountNumber"])
	balance, _ := service.GetBalance()
	respondWithJSON(w, 200, balance)
}

func (handler accountHandlerImpl) GetTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := handler.getTransactionService()
	transactions, _ := service.GetTransactions(vars["AccountNumber"])
	respondWithJSON(w, 200, transactions)
}

func (handler accountHandlerImpl) TransferMoney(w http.ResponseWriter, r *http.Request) {
	var transferRequest dto.TransferRequest

	err := json.NewDecoder(r.Body).Decode(&transferRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	service := handler.getAccountService(transferRequest.FromAccount)
	trxService := handler.getTransactionService()

	balance := ExecuteTransfer(service, trxService, transferRequest)
	respondWithJSON(w, 200, balance)

}

func (handler accountHandlerImpl) Deposit(w http.ResponseWriter, r *http.Request) {
	var depositRequest dto.DepositRequest
	err := json.NewDecoder(r.Body).Decode(&depositRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	service := handler.getAccountService(depositRequest.ToAccount)
	trxService := handler.getTransactionService()
	balance := ExecuteDeposit(service, trxService, depositRequest)
	respondWithJSON(w, 200, balance)
}

func (a accountHandlerImpl) getAccountService(accountNumber string) port.AccountService {
	return services.NewAccountService(accountNumber, a.AccountRepository)

}

func (a accountHandlerImpl) getTransactionService() port.TransactionService {
	return services.NewTransactionService(a.AccountRepository)

}

//TODO MOVE THIS TO AN UTIL module
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(dto.ResponseDto{Data: payload})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ExecuteDeposit(accountService port.AccountService, transactionService port.TransactionService,
	depositRequest dto.DepositRequest) float32 {

	accountService.Deposit(depositRequest.Amount)
	transactionService.SaveTransaction(dto.TransactionDto{
		AccountTo: depositRequest.ToAccount,
		Amount:    depositRequest.Amount,
		Type:      port.DepositType,
	})
	balance, _ := accountService.GetBalance()
	return balance
}

func ExecuteTransfer(accountService port.AccountService, transactionService port.TransactionService,
	transferRequest dto.TransferRequest) float32 {
	err := accountService.TransferMoneyTo(transferRequest.ToAccount, transferRequest.Amount)
	if err == nil {
		transactionService.SaveTransaction(dto.TransactionDto{
			AccountFrom: transferRequest.FromAccount,
			AccountTo:   transferRequest.ToAccount,
			Amount:      transferRequest.Amount,
			Type:        port.TransferType,
		})
	}

	balance, _ := accountService.GetBalance()
	return balance
}

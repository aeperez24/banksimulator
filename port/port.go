package port

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"net/http"
)

type AccountService interface {
	GetBalance() (float32, error)
	TransferMoneyTo(accountNumber string, amount float32) error
	Deposit(amount float32) error
}

type TransactionService interface {
	GetTransactions(string) ([]model.Transaction, error)
	SaveTransaction(dto.TransactionDto) error
}

type AccountHandler interface {
	GetBalance(http.ResponseWriter, *http.Request)
	TransferMoney(http.ResponseWriter, *http.Request)
	Deposit(http.ResponseWriter, *http.Request)
}
type Server interface {
	Start()
	Stop()
}

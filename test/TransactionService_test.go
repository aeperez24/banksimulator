package test

import (
	"aeperez24/banksimulator/services"
)

func TestGetTransactions() {
	repo := AccountRepositoryMock{}
	//TODO config mock
	accountNumber := ""
	service := services.NewTransactionService(accountNumber, repo)
	transactions, _ := service.GetTransactions()
	//TODO assert on transactions
}

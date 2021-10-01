package test

import (
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/services"
	"testing"
)

type AccountRepositoryMock struct {
	FindAccountByAccountNumberFn func(account string) model.Account
	ModifyBalanceForAccountFn    func(accountNumber string, amount float32) bool
}

func (a AccountRepositoryMock) FindAccountByAccountNumber(account string) model.Account {
	return a.FindAccountByAccountNumberFn(account)
}

func (a AccountRepositoryMock) ModifyBalanceForAccount(accountNumber string, amount float32) bool {
	return a.ModifyBalanceForAccountFn(accountNumber, amount)
}

func getAccountRepositoryMock() model.AccountRepository {
	accountsMap := make(map[string]model.Account)
	accountsMap["1"] = model.Account{AccountNumber: "1", Balance: 100}
	accountsMap["2"] = model.Account{AccountNumber: "2", Balance: 100}
	findAccountByAccountNumber := func(number string) model.Account {
		return accountsMap[number]
	}
	return AccountRepositoryMock{FindAccountByAccountNumberFn: findAccountByAccountNumber}
}

func TestGetBalance(t *testing.T) {
	mock := getAccountRepositoryMock()
	service := services.NewAccountService("1", mock)
	result := service.GetBalance()
	expected := 100
	if 100 != result {
		t.Errorf("expected %v and received %v", expected, result)
	}

}

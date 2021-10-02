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
	accountsMap["2"] = model.Account{AccountNumber: "2", Balance: 200}
	findAccountByAccountNumber := func(number string) model.Account {
		return accountsMap[number]
	}
	modifyBalance := func(accountNumber string, amount float32) bool {

		account := accountsMap[accountNumber]
		account.Balance = account.Balance + amount
		accountsMap[accountNumber] = account

		return true
	}
	return AccountRepositoryMock{FindAccountByAccountNumberFn: findAccountByAccountNumber,
		ModifyBalanceForAccountFn: modifyBalance}
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

func TestMustTransferAmountSuccesfully(t *testing.T) {
	mock := getAccountRepositoryMock()
	service := services.NewAccountService("1", mock)
	result := service.TransferMoneyTo("2", 50)
	balanceAccount1 := mock.FindAccountByAccountNumber("1").Balance
	balanceAccount2 := mock.FindAccountByAccountNumber("2").Balance
	expectedBalanceAccount1 := float32(50)
	expectedBalanceAccount2 := float32(250)
	expectedResult := true
	if expectedBalanceAccount1 != balanceAccount1 {
		t.Errorf("expected %v and received %v", expectedBalanceAccount1, balanceAccount1)
	}
	if expectedBalanceAccount2 != balanceAccount2 {
		t.Errorf("expected %v and received %v", expectedBalanceAccount2, balanceAccount2)
	}

	if expectedResult != result {
		t.Errorf("expected %v and received %v", expectedResult, result)
	}

}

func TestMustNotTransferAmountWhenIsGreaterThanBalance(t *testing.T) {
	mock := getAccountRepositoryMock()
	service := services.NewAccountService("1", mock)
	result := service.TransferMoneyTo("2", 500)
	balanceAccount1 := mock.FindAccountByAccountNumber("1").Balance
	balanceAccount2 := mock.FindAccountByAccountNumber("2").Balance
	expectedBalanceAccount1 := float32(100)
	expectedBalanceAccount2 := float32(200)
	expectedResult := false

	if expectedBalanceAccount1 != balanceAccount1 {
		t.Errorf("expected %v and received %v", expectedBalanceAccount1, balanceAccount1)
	}
	if expectedBalanceAccount2 != balanceAccount2 {
		t.Errorf("expected %v and received %v", expectedBalanceAccount2, balanceAccount2)
	}

	if expectedResult != result {
		t.Errorf("expected %v and received %v", expectedResult, result)
	}

}

func TestDepositBalanceSuccesfully(t *testing.T) {
	mock := getAccountRepositoryMock()
	service := services.NewAccountService("1", mock)
	expectedBalanceAccount := float32(180)
	service.Deposit(80)
	balanceAccount := mock.FindAccountByAccountNumber("1").Balance

	if expectedBalanceAccount != balanceAccount {
		t.Errorf("expected %v and received %v", expectedBalanceAccount, balanceAccount)
	}

}

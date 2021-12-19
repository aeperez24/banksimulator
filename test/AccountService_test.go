package test

import (
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/services"
	"testing"
)

func getAccountRepositoryMock() model.AccountRepository {
	accountsMap := make(map[string]model.Account)
	accountsMap["1"] = model.Account{AccountNumber: "1", Balance: 100}
	accountsMap["2"] = model.Account{AccountNumber: "2", Balance: 200}
	findAccountByAccountNumber := func(number string) model.Account {
		return accountsMap[number]
	}
	modifyBalance := func(accountNumber string, amount float32) error {

		account := accountsMap[accountNumber]
		account.Balance = account.Balance + amount
		accountsMap[accountNumber] = account

		return nil
	}
	return AccountRepositoryMock{FindAccountByAccountNumberFn: findAccountByAccountNumber,
		ModifyBalanceForAccountFn: modifyBalance}
}

func TestGetBalance(t *testing.T) {
	mock := getAccountRepositoryMock()
	service := services.NewAccountService("1", mock)
	result, _ := service.GetBalance()
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
	if expectedBalanceAccount1 != balanceAccount1 {
		t.Errorf("expected %v and received %v", expectedBalanceAccount1, balanceAccount1)
	}
	if expectedBalanceAccount2 != balanceAccount2 {
		t.Errorf("expected %v and received %v", expectedBalanceAccount2, balanceAccount2)
	}
	if result != nil {
		t.Errorf("expected %v and received %v", nil, result)
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

	if expectedBalanceAccount1 != balanceAccount1 {
		t.Errorf("expected %v and received %v", expectedBalanceAccount1, balanceAccount1)
	}
	if expectedBalanceAccount2 != balanceAccount2 {
		t.Errorf("expected %v and received %v", expectedBalanceAccount2, balanceAccount2)
	}

	if nil == result {
		t.Errorf("expected error")
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

func TestWithdrawBalanceSuccesfully(t *testing.T) {
	mock := getAccountRepositoryMock()
	service := services.NewAccountService("1", mock)
	expectedBalanceAccount := float32(20)
	service.Withdraw(80)
	balanceAccount := mock.FindAccountByAccountNumber("1").Balance

	if expectedBalanceAccount != balanceAccount {
		t.Errorf("expected %v and received %v", expectedBalanceAccount, balanceAccount)
	}

}
func TestWithdrawGreaterThanBalanceError(t *testing.T) {
	mock := getAccountRepositoryMock()
	service := services.NewAccountService("1", mock)
	resultError := service.Withdraw(150)
	if resultError == nil {
		t.Error("error is expected")
	}

}

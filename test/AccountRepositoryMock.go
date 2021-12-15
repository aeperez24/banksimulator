package test

import "aeperez24/banksimulator/model"

type AccountRepositoryMock struct {
	FindAccountByAccountNumberFn func(account string) model.Account
	ModifyBalanceForAccountFn    func(accountNumber string, amount float32) error
	SaveTransactionFn            func(account string, transaction model.Transaction) error
}

func (a AccountRepositoryMock) FindAccountByAccountNumber(account string) model.Account {
	return a.FindAccountByAccountNumberFn(account)
}

func (a AccountRepositoryMock) ModifyBalanceForAccount(accountNumber string, amount float32) error {
	return a.ModifyBalanceForAccountFn(accountNumber, amount)
}

func (a AccountRepositoryMock) SaveTransaction(account string, transaction model.Transaction) error {
	return a.SaveTransactionFn(account, transaction)
}

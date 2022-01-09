package test

import "aeperez24/banksimulator/model"

type AccountRepositoryMock struct {
	FindAccountByAccountNumberFn func(account string) model.Account
	ModifyBalanceForAccountFn    func(accountNumber string, amount float32) error
	SaveTransactionFn            func(account string, transaction model.Transaction) error
	CreateAccountFn              func(account model.Account) (interface{}, error)
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

func (a AccountRepositoryMock) CreateAccount(account model.Account) (interface{}, error) {
	return a.CreateAccountFn(account)
}

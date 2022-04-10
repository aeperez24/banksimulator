package test

import (
	"aeperez24/banksimulator/model"

	"github.com/stretchr/testify/mock"
)

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

type AccountRepositoryMockTestify struct {
	mock.Mock
}

func (m AccountRepositoryMockTestify) FindAccountByAccountNumber(account string) model.Account {
	args := m.Called(account)
	return args.Get(0).(model.Account)
}

func (m AccountRepositoryMockTestify) ModifyBalanceForAccount(accountNumber string, amount float32) error {
	args := m.Called(accountNumber, amount)
	return args.Error(0)
}

func (m AccountRepositoryMockTestify) SaveTransaction(account string, transaction model.Transaction) error {
	args := m.Called(account, transaction)
	return args.Error(0)
}

func (m AccountRepositoryMockTestify) CreateAccount(account model.Account) (interface{}, error) {
	args := m.Called(account)
	return args.Get(0), args.Error(0)
}

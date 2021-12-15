package services

import "aeperez24/banksimulator/model"

type TransactionService interface {
	GetTransactions() ([]model.Transaction, error)
	SaveTransaction(model.Transaction) error
}

type transactionServiceImpl struct {
	AccountNumber     string
	AccountRepository model.AccountRepository
}

func (service transactionServiceImpl) GetTransactions() ([]model.Transaction, error) {
	return []model.Transaction{}, nil
}

func (service transactionServiceImpl) SaveTransaction(model.Transaction) error {
	return nil
}

func NewTransactionService(accountNumber string, accountRepository model.AccountRepository) TransactionService {
	return transactionServiceImpl{accountNumber, accountRepository}
}

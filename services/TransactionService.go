package services

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
)

type transactionServiceImpl struct {
	AccountRepository model.AccountRepository
}

func (service transactionServiceImpl) GetTransactions(accountNumber string) ([]model.Transaction, error) {
	account := service.AccountRepository.FindAccountByAccountNumber(accountNumber)
	return account.Transactions, nil
}

func (service transactionServiceImpl) SaveTransaction(transactiondto dto.TransactionDto) error {
	transaction := model.Transaction{
		AccountFrom: transactiondto.AccountFrom,
		AccountTo:   transactiondto.AccountTo,
		Amount:      transactiondto.Amount,
	}

	service.AccountRepository.SaveTransaction(transaction.AccountFrom, transaction)
	service.AccountRepository.SaveTransaction(transaction.AccountTo, transaction)

	return nil
}

func NewTransactionService(accountRepository model.AccountRepository) port.TransactionService {
	return transactionServiceImpl{accountRepository}
}

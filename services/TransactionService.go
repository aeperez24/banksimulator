package services

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"errors"
)

type transactionServiceImpl struct {
	AccountRepository model.AccountRepository
	saveFuncMap       map[model.TransactionType]func(transactionServiceImpl, dto.TransactionDto) error
}

func (service transactionServiceImpl) GetTransactions(accountNumber string) ([]model.Transaction, error) {
	account := service.AccountRepository.FindAccountByAccountNumber(accountNumber)
	return account.Transactions, nil
}

func (service transactionServiceImpl) SaveTransaction(transactiondto dto.TransactionDto) error {
	transactionType := model.TransactionType(transactiondto.Type)
	selectedSaveFunction := service.saveFuncMap[transactionType]
	if selectedSaveFunction != nil {
		return selectedSaveFunction(service, transactiondto)

	} else {
		return errors.New("invalid transaction type")
	}
}

func NewTransactionService(accountRepository model.AccountRepository) port.TransactionService {
	saveFuncMap := make(map[model.TransactionType]func(transactionServiceImpl, dto.TransactionDto) error)
	saveFuncMap[model.TransferType] = saveTransferTransaction
	saveFuncMap[model.DepositType] = saveDepositTransaction
	saveFuncMap[model.WithdrawType] = saveWithdrawTransaction
	return transactionServiceImpl{accountRepository, saveFuncMap}
}

func saveTransferTransaction(service transactionServiceImpl, transactiondto dto.TransactionDto) error {
	transaction := model.Transaction{
		AccountFrom: transactiondto.AccountFrom,
		AccountTo:   transactiondto.AccountTo,
		Amount:      transactiondto.Amount,
		Type:        model.TransactionType(transactiondto.Type),
	}

	service.AccountRepository.SaveTransaction(transaction.AccountFrom, transaction)
	service.AccountRepository.SaveTransaction(transaction.AccountTo, transaction)

	return nil
}

func saveDepositTransaction(service transactionServiceImpl, transactiondto dto.TransactionDto) error {
	transaction := model.Transaction{
		AccountTo: transactiondto.AccountTo,
		Amount:    transactiondto.Amount,
		Type:      model.TransactionType(transactiondto.Type),
	}

	service.AccountRepository.SaveTransaction(transaction.AccountTo, transaction)

	return nil
}

func saveWithdrawTransaction(service transactionServiceImpl, transactiondto dto.TransactionDto) error {
	transaction := model.Transaction{
		AccountFrom: transactiondto.AccountFrom,
		Amount:      transactiondto.Amount,
		Type:        model.TransactionType(transactiondto.Type),
	}

	service.AccountRepository.SaveTransaction(transaction.AccountFrom, transaction)

	return nil
}

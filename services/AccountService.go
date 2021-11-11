package services

import (
	"aeperez24/banksimulator/model"
)

type AccountService interface {
	GetBalance() float32
	TransferMoneyTo(accountNumber string, amount float32) bool
	Deposit(amount float32) bool
}

type accountServiceImp struct {
	AccountNumber     string
	AccountRepository model.AccountRepository
}

func NewAccountService(accountNumber string, accountRepository model.AccountRepository) AccountService {
	return accountServiceImp{accountNumber, accountRepository}
}

func (acountService accountServiceImp) GetBalance() float32 {
	return acountService.AccountRepository.
		FindAccountByAccountNumber(acountService.AccountNumber).Balance
}

func (accountService accountServiceImp) TransferMoneyTo(toAccountNumber string, amount float32) bool {

	if amount <= accountService.GetBalance() {

		repository := accountService.AccountRepository
		repository.ModifyBalanceForAccount(accountService.AccountNumber, -amount)
		repository.ModifyBalanceForAccount(toAccountNumber, amount)
		return true

	}
	return false
}

func (accountService accountServiceImp) Deposit(amount float32) bool {
	repository := accountService.AccountRepository
	repository.ModifyBalanceForAccount(accountService.AccountNumber, amount)
	return true
}

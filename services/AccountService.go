package services

import (
	"aeperez24/banksimulator/model"
)

type AccountService interface {
	GetBalance() float32
	TransferMoneyTo(accountNumber string) bool
}

type accountServiceImp struct {
	AcountNumber     string
	AcountRepository model.AccountRepository
}

func NewAccountService(accountNumber string, accountRepository model.AccountRepository) AccountService {
	return accountServiceImp{accountNumber, accountRepository}
}

func (acountService accountServiceImp) GetBalance() float32 {
	return acountService.AcountRepository.
		FindAccountByAccountNumber(acountService.AcountNumber).Balance
}

func (acountService accountServiceImp) TransferMoneyTo(accountNumber string) bool {
	return false
}

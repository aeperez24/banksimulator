package services

import (
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"errors"
	"log"
)

type accountServiceImp struct {
	AccountNumber     string
	AccountRepository model.AccountRepository
}

func NewAccountServiceProvider(AccountRepository model.AccountRepository) func(string) port.AccountService {
	result := func(accountNumber string) port.AccountService {
		return NewAccountService(accountNumber, AccountRepository)
	}
	return result
}
func NewAccountService(accountNumber string, accountRepository model.AccountRepository) port.AccountService {
	return accountServiceImp{accountNumber, accountRepository}
}

func (acountService accountServiceImp) GetBalance() (float32, error) {
	return acountService.AccountRepository.
		FindAccountByAccountNumber(acountService.AccountNumber).Balance, nil
}

func (accountService accountServiceImp) TransferMoneyTo(toAccountNumber string, amount float32) error {
	balance, error := accountService.GetBalance()
	if error != nil {
		log.Fatal(error)
		return error
	}

	if amount <= balance {

		repository := accountService.AccountRepository
		repository.ModifyBalanceForAccount(accountService.AccountNumber, -amount)
		repository.ModifyBalanceForAccount(toAccountNumber, amount)
		return nil

	} else {
		return errors.New("amount higher than banlance")
	}

}

func (accountService accountServiceImp) Deposit(amount float32) error {
	repository := accountService.AccountRepository
	return repository.ModifyBalanceForAccount(accountService.AccountNumber, amount)
}
func (accountService accountServiceImp) Withdraw(amount float32) error {
	balance, _ := accountService.GetBalance()
	if balance > amount {
		return accountService.Deposit(-amount)
	} else {
		return errors.New("amount is greater than current balance")
	}

}

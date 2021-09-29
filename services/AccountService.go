package services

import (
	"aeperez24/banksimulator/model"
)

type AccountService interface {
	getBalance()
	transferMoneyTo()
}

type AccountServiceImp struct {
	Acount           string
	AcountRepository model.AccountRepository
}

func newAccountService(acount model.Account, AcountRepository model.AccountRepository) AccountService {
	return AccountServiceImp{}
}

func (acountService AccountServiceImp) getBalance() {

}

func (acountService AccountServiceImp) transferMoneyTo() {

}

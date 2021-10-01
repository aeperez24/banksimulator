package model

type AccountRepository interface {
	FindAccountByAccountNumber(account string) Account
	ModifyBalanceForAccount(accountNumber string, amount float32) bool
}

package model

import "time"

type Transaction struct {
	AccountFrom string
	AccountTo   string
	Amount      float32
	Date        time.Time
	Type        TransactionType
}
type TransactionType string

const (
	TransferType TransactionType = "Transfer"
	DepositType  TransactionType = "Deposit"
	WithdrawType TransactionType = "Withdraw"
)

type Account struct {
	AccountNumber string
	Balance       float32
	Transactions  []Transaction
}

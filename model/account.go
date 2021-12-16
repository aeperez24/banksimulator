package model

import "time"

type Transaction struct {
	AccountFrom string
	AccountTo   string
	Amount      float32
	Date        time.Time
}

type Account struct {
	AccountNumber string
	Balance       float32
	Transactions  []Transaction
}

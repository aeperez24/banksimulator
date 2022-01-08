package dto

import "time"

type TransferRequest struct {
	FromAccount string
	ToAccount   string
	Amount      float32
}

type DepositRequest struct {
	ToAccount string
	Amount    float32
}

type TransactionDto struct {
	AccountFrom string
	AccountTo   string
	Amount      float32
	Date        time.Time
	Type        string
}

type UserWithPasswordDto struct {
	BasicUserDto
	Password string
}
type BasicUserDto struct {
	Username   string
	IDDocument string
}
type ResponseDto struct {
	Data interface{}
}

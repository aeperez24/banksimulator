package dto

type TransferRequest struct {
	FromAccount string
	ToAccount   string
	Amount      float32
}

type DepositRequest struct {
	ToAccount string
	Amount    float32
}

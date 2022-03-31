package usercase

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"aeperez24/banksimulator/services"
	"context"
	"reflect"
)

type UserCaseError struct {
	error
	Code    int
	Message string
}

//USER
type UserUsercase struct {
	AccountRepository model.AccountRepository
	UserService       port.UserService
}

func (userUserCase UserUsercase) CreateUser(user dto.UserWithPasswordDto) UserCaseError {
	basicUser := userUserCase.UserService.FindBasicUser(user.Username)
	if basicUser != (dto.BasicUserDto{}) {
		return UserCaseError{nil, 400, "username already exists"}
	}
	userByDocument := userUserCase.UserService.FindBasicUserByDocument(user.Username)
	if userByDocument != (dto.BasicUserDto{}) {
		return UserCaseError{nil, 400, "DocumentId already exists already exists"}

	}

	acc := userUserCase.AccountRepository.FindAccountByAccountNumber(user.IDDocument)

	if !reflect.DeepEqual(acc, model.Account{}) {
		return UserCaseError{nil, 400, "account already exists"}
	}

	_, err := userUserCase.AccountRepository.CreateAccount(model.Account{AccountNumber: user.IDDocument})

	if err != nil {
		return UserCaseError{err, 500, "error creating account"}
	}
	err = userUserCase.UserService.CreateUser(user)

	if err != nil {
		return UserCaseError{err, 500, "error creating user"}

	}
	return UserCaseError{}
}

//ACCOUNT
type AccountUsercase struct {
	AccountRepository  model.AccountRepository
	TransactionService port.TransactionService
}

func (acoountUseCase AccountUsercase) GetBalance(accountNumber string) (float32, UserCaseError) {
	service := acoountUseCase.getAccountService(accountNumber)
	balance, _ := service.GetBalance()
	return balance, UserCaseError{}
}

func (useCase AccountUsercase) GetTransactions(ctx context.Context, accountNumber string) ([]model.Transaction, UserCaseError) {
	user := (ctx.Value(port.LoggedUserKey)).(dto.BasicUserDto)
	if user.IDDocument != accountNumber {
		return nil, UserCaseError{nil, 403, ""}

	}
	transactions, err := useCase.TransactionService.GetTransactions(accountNumber)
	if err != nil {
		return nil, UserCaseError{err, 500, "Internal Error"}

	}
	return transactions, UserCaseError{}

}

func (a AccountUsercase) getAccountService(accountNumber string) port.AccountService {
	return services.NewAccountService(accountNumber, a.AccountRepository)

}

func (useCase AccountUsercase) TransferMoney(ctx context.Context, transferRequest dto.TransferRequest) (float32, UserCaseError) {
	user := (ctx.Value(port.LoggedUserKey)).(dto.BasicUserDto)
	if user.IDDocument != transferRequest.FromAccount {
		return 0, UserCaseError{nil, 403, ""}

	}
	service := useCase.getAccountService(transferRequest.FromAccount)

	err := service.TransferMoneyTo(transferRequest.ToAccount, transferRequest.Amount)
	if err == nil {
		useCase.TransactionService.SaveTransaction(dto.TransactionDto{
			AccountFrom: transferRequest.FromAccount,
			AccountTo:   transferRequest.ToAccount,
			Amount:      transferRequest.Amount,
			Type:        port.TransferType,
		})
	}

	balance, err := service.GetBalance()

	if err != nil {
		return 0, UserCaseError{err, 500, ""}
	}
	return balance, UserCaseError{}

}

func (useCase AccountUsercase) Deposit(ctx context.Context, depositRequest dto.DepositRequest) (float32, UserCaseError) {
	user := (ctx.Value(port.LoggedUserKey)).(dto.BasicUserDto)
	if user.IDDocument != depositRequest.ToAccount {
		return 0, UserCaseError{nil, 403, ""}

	}
	service := useCase.getAccountService(depositRequest.ToAccount)
	service.Deposit(depositRequest.Amount)
	useCase.TransactionService.SaveTransaction(dto.TransactionDto{
		AccountTo: depositRequest.ToAccount,
		Amount:    depositRequest.Amount,
		Type:      port.DepositType,
	})
	balance, err := service.GetBalance()

	if err != nil {
		return 0, UserCaseError{err, 500, ""}
	}
	return balance, UserCaseError{}

}

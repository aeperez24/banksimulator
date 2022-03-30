package usercase

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"reflect"
)

type UserCaseError struct {
	error
	Code    int
	Message string
}
type AccountUsercase struct {
}

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

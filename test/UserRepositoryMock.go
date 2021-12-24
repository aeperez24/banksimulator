package test

import "aeperez24/banksimulator/model"

type UserRepositoryMock struct {
	FindUserByNameFn       func(username string) model.User
	CreateUserFn           func(user model.User) (interface{}, error)
	FindUserByIdDocumentFn func(idDocument string) model.User
}

func (a UserRepositoryMock) FindUserByName(username string) model.User {
	return a.FindUserByNameFn(username)
}

func (a UserRepositoryMock) CreateUser(user model.User) (interface{}, error) {
	return a.CreateUserFn(user)
}

func (a UserRepositoryMock) FindUserByIdDocument(idDocument string) model.User {
	return a.FindUserByIdDocumentFn(idDocument)
}

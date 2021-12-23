package test

import "aeperez24/banksimulator/model"

type UserRepositoryMock struct {
	FindUserByNameFn func(username string) model.User
	CreateUserFn     func(user model.User) (interface{}, error)
}

func (a UserRepositoryMock) FindUserByName(username string) model.User {
	return a.FindUserByNameFn(username)
}

func (a UserRepositoryMock) CreateUser(user model.User) (interface{}, error) {
	return a.CreateUserFn(user)
}

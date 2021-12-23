package services

import (
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"

	"errors"
)

type userServiceImpl struct {
	UserRepository model.UserRepository
}

func (userService userServiceImpl) CreateUser(user model.User) error {
	foundUser := userService.UserRepository.FindUserByName(user.Username)

	if (foundUser != model.User{}) {
		return errors.New("username already exists")
	}

	_, err := userService.UserRepository.CreateUser(user)
	return err
}

func (userService userServiceImpl) ValidateUserameAndPassword(username string, password string) bool {
	return false
}

func NewUserService(repo model.UserRepository) port.UserService {
	return userServiceImpl{
		UserRepository: repo,
	}

}

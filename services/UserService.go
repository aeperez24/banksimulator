package services

import (
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"crypto/sha256"
	"errors"
	"fmt"
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

func (userService userServiceImpl) FindUser(username string) model.User {
	return userService.UserRepository.FindUserByName(username)
}
func (userService userServiceImpl) ValidateUserameAndPassword(username string, password string) bool {
	sha := sha256.Sum256([]byte(password))
	user := userService.FindUser(username)
	strSha := fmt.Sprintf("%x", sha)
	return strSha == user.Password
}

func NewUserService(repo model.UserRepository) port.UserService {
	return userServiceImpl{
		UserRepository: repo,
	}

}

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
	if user.Password == "" {
		return errors.New("password is required")
	}
	if user.Username == "" {
		return errors.New("username is required")
	}

	if user.IDDocument == "" {
		return errors.New("idDocument is required")
	}

	foundUser := userService.UserRepository.FindUserByName(user.Username)
	if (foundUser != model.User{}) {
		return errors.New("username already exists")
	}
	idDocumentAlreadyRegistered := userService.UserRepository.FindUserByIdDocument(user.IDDocument) != model.User{}
	if idDocumentAlreadyRegistered {
		return errors.New("document already registered")
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

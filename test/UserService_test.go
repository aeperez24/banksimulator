package test

import (
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/services"
	"testing"
)

func TestCreateUser(t *testing.T) {
	createCalled := false
	repo := UserRepositoryMock{FindUserByNameFn: func(username string) model.User {
		if username != "user" {
			return model.User{Username: "user", Password: "pass"}
		} else {
			return model.User{}
		}
	}, CreateUserFn: func(user model.User) (interface{}, error) {
		createCalled = true
		return "id", nil
	}}
	service := services.NewUserService(repo)
	user := model.User{
		Username: "user",
	}
	error := service.CreateUser(user)
	if error != nil {
		t.Errorf("expected %v and received %v", nil, error)
	}
	if !createCalled {
		t.Errorf("expected %v and received %v", true, createCalled)
	}

}

func TestMustFailWhenCreateUserWithUsernameInDatabase(t *testing.T) {
	repo := UserRepositoryMock{FindUserByNameFn: func(username string) model.User {
		if username == "user" {
			return model.User{Username: "user", Password: "pass"}
		} else {
			return model.User{}
		}
	}, CreateUserFn: func(user model.User) (interface{}, error) {
		return "id", nil
	}}
	service := services.NewUserService(repo)
	user := model.User{
		Username: "user",
	}
	error := service.CreateUser(user)
	if error == nil {
		t.Errorf("expected error")
	}
}

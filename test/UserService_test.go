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
		Username:   "user",
		Password:   "pass",
		IDDocument: "document",
	}
	error := service.CreateUser(user)
	if error == nil {
		t.Errorf("expected error")
	}
}

func TestValidatePasswordSuccess(t *testing.T) {
	passwordSha := "d74ff0ee8da3b9806b18c877dbf29bbde50b5bd8e4dad7a3a725000feb82e8f1"

	repo := UserRepositoryMock{FindUserByNameFn: func(username string) model.User {
		if username == "user" {
			return model.User{Username: "user", Password: passwordSha}
		} else {
			return model.User{}
		}
	}, CreateUserFn: func(user model.User) (interface{}, error) {
		return "id", nil
	}}
	service := services.NewUserService(repo)

	isValid := service.ValidateUserameAndPassword("user", "pass")
	if !isValid {
		t.Error("expected valid")
	}
}

func TestValidatePasswordFail(t *testing.T) {
	passwordSha := "d74ff0ee8da3b9806b18c877dbf29bbde50b5bd8e4dad7a3a725000feb82e8f1"

	repo := UserRepositoryMock{FindUserByNameFn: func(username string) model.User {
		if username == "user" {
			return model.User{Username: "user", Password: passwordSha}
		} else {
			return model.User{}
		}
	}, CreateUserFn: func(user model.User) (interface{}, error) {
		return "id", nil
	}}
	service := services.NewUserService(repo)

	isValid := service.ValidateUserameAndPassword("user", "badpass")
	if isValid {
		t.Error("expected not valid")
	}
}

func TestMustFailWhenCreateUserWithoutDocumentId(t *testing.T) {
	repo := UserRepositoryMock{FindUserByNameFn: func(username string) model.User {
		if username != "user" {
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
		Password: "pass",
	}
	error := service.CreateUser(user)
	if error == nil {
		t.Errorf("expected error")
	}
}

func TestMustFailWhenCreateUserWithoutPasword(t *testing.T) {
	repo := UserRepositoryMock{FindUserByNameFn: func(username string) model.User {
		if username != "user" {
			return model.User{Username: "user", Password: "pass"}
		} else {
			return model.User{}
		}
	}, CreateUserFn: func(user model.User) (interface{}, error) {
		return "id", nil
	}}
	service := services.NewUserService(repo)
	user := model.User{
		Username:   "user",
		IDDocument: "document",
	}
	error := service.CreateUser(user)
	if error == nil {
		t.Errorf("expected error")
	}
}

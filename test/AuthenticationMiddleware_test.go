package test

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/middleware"
	"aeperez24/banksimulator/port"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type tokenServiceMock struct {
	CreateTokenFn                 func(dto.BasicUserDto) (string, error)
	ExtractBasicUseDtoFromTokenFn func(receivedToken string) (dto.BasicUserDto, error)
}

func (mock tokenServiceMock) CreateToken(in dto.BasicUserDto) (string, error) {
	return mock.CreateTokenFn(in)
}
func (mock tokenServiceMock) ExtractBasicUseDtoFromToken(receivedToken string) (dto.BasicUserDto, error) {
	return mock.ExtractBasicUseDtoFromTokenFn(receivedToken)
}

func TestAuthenticateSuccessfully(t *testing.T) {
	executed := false
	extractMock := func(receivedToken string) (dto.BasicUserDto, error) {

		return dto.BasicUserDto{}, nil
	}
	mid := middleware.NewAuthenticationMiddlewre(tokenServiceMock{ExtractBasicUseDtoFromTokenFn: extractMock})
	handler := func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(port.LoggedUserKey)
		if user == nil {
			t.Error("user expected")
		}
		executed = true
	}
	filter := mid.Filter(handler)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	writer := httptest.NewRecorder()
	req.Header.Add("Authorization", "bearer 123")
	filter(writer, req)
	if !executed {
		t.Error("handler not executed")
	}
}

func TestAuthenticateRejectWhenNotToken(t *testing.T) {
	executed := false
	extractMock := func(receivedToken string) (dto.BasicUserDto, error) {

		return dto.BasicUserDto{}, nil
	}
	mid := middleware.NewAuthenticationMiddlewre(tokenServiceMock{ExtractBasicUseDtoFromTokenFn: extractMock})
	handler := func(w http.ResponseWriter, r *http.Request) {
		executed = true
	}
	filter := mid.Filter(handler)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	writer := httptest.NewRecorder()
	filter(writer, req)
	if executed {
		t.Error("handler shoulnot be executed")
	}
}

func TestAuthenticateRejectWhenTokenServiceReturnError(t *testing.T) {
	executed := false
	extractMock := func(receivedToken string) (dto.BasicUserDto, error) {

		return dto.BasicUserDto{}, errors.New("")
	}
	mid := middleware.NewAuthenticationMiddlewre(tokenServiceMock{ExtractBasicUseDtoFromTokenFn: extractMock})
	handler := func(w http.ResponseWriter, r *http.Request) {
		executed = true
	}
	filter := mid.Filter(handler)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	writer := httptest.NewRecorder()
	filter(writer, req)
	if executed {
		t.Error("handler should not be executed")
	}
}

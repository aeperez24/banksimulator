package test

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/services"
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "123"

func TestCreateToken(t *testing.T) {
	user := dto.BasicUserDto{
		Username:   "user",
		IDDocument: "id",
	}
	tokenService := services.NewTokenService(secretKey)
	result, err := tokenService.CreateToken(user)

	if err != nil {
		t.Errorf("%v", err)
	}

	userFromToken, err := getDtoFromToken(result)
	if err != nil {
		t.Errorf("%v", err)
	}

	if user != userFromToken {
		t.Errorf("expected %v and received %v", user, userFromToken)
	}

}

func TestExtractUserFromTokenToken(t *testing.T) {
	user := dto.BasicUserDto{
		Username:   "user",
		IDDocument: "id",
	}
	tokenService := services.NewTokenService(secretKey)
	token, _ := tokenService.CreateToken(user)
	userFromToken, err := tokenService.ExtractBasicUseDtoFromToken(token)

	if err != nil {
		t.Errorf("%v", err)
	}

	fmt.Printf("%v", userFromToken)
	if user != userFromToken {
		t.Errorf("expected %v and received %v", user, userFromToken)
	}

}

func getDtoFromToken(receivedToken string) (dto.BasicUserDto, error) {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	result := dto.BasicUserDto{
		Username:   claims["user_name"].(string),
		IDDocument: claims["document_id"].(string),
	}
	return result, err
}

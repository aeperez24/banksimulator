package services

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/port"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenServiceImpl struct {
	accessSecret string
}

func (service tokenServiceImpl) CreateToken(user dto.BasicUserDto) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_name"] = user.Username
	atClaims["document_id"] = user.IDDocument
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(service.accessSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service tokenServiceImpl) ExtractBasicUseDtoFromToken(receivedToken string) (dto.BasicUserDto, error) {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(service.accessSecret), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	result := dto.BasicUserDto{
		Username:   claims["user_name"].(string),
		IDDocument: claims["document_id"].(string),
	}
	return result, err
}

func NewTokenService(accessSecret string) port.TokenService {
	return tokenServiceImpl{
		accessSecret: accessSecret,
	}
}

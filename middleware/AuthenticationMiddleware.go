package middleware

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/port"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
)

type authenticationMiddleware struct {
	tokenService port.TokenService
}

func (middleware authenticationMiddleware) Filter(handler HttpHandler) HttpHandler {

	return func(w http.ResponseWriter, r *http.Request) {
		user, err := middleware.extractToken(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			return
		}
		log.Printf("user :%v", user)
		handler(w, r.WithContext(context.WithValue(r.Context(), port.LoggedUserKey, user)))
	}

}

func (middleware authenticationMiddleware) extractToken(r *http.Request) (dto.BasicUserDto, error) {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		log.Println("Authorization header not found")
		return dto.BasicUserDto{}, errors.New("invalid token")
	}
	return middleware.tokenService.ExtractBasicUseDtoFromToken(strArr[1])
}

func NewAuthenticationMiddlware(tokenService port.TokenService) Middleware {
	return authenticationMiddleware{tokenService}

}

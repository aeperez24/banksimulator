package middleware

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/port"
	"context"
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type authenticationMiddlewareGin struct {
	tokenService port.TokenService
}

func (middleware authenticationMiddlewareGin) getMiddleware() func(c *gin.Context) {

	return func(c *gin.Context) {
		user, err := middleware.extractToken(c)
		if err != nil {
			c.JSON(400, nil)
			return
		}
		log.Printf("user :%v", user)
		ncontext := context.WithValue(c.Request.Context(), port.LoggedUserKey, user)
		c.Request = c.Request.WithContext(ncontext)

		c.Next()
	}

}

func (middleware authenticationMiddlewareGin) extractToken(c *gin.Context) (dto.BasicUserDto, error) {
	bearToken := c.Request.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		log.Println("Authorization header not found")
		return dto.BasicUserDto{}, errors.New("invalid token")
	}
	return middleware.tokenService.ExtractBasicUseDtoFromToken(strArr[1])
}

func NewAuthenticationMiddlwareGin(tokenService port.TokenService) func(c *gin.Context) {
	return authenticationMiddlewareGin{tokenService}.getMiddleware()

}

package handler

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/middleware"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/services"
	"aeperez24/banksimulator/usercase"
)

func BuildServerConfigGin(portNumber string, tokenKey string, mongo config.MongoCofig) ServerConfigurationGin {
	repo := model.NewAccountMongoRepository(mongo)
	tokenService := services.NewTokenService(tokenKey)
	userRepo := model.NewUserMongoRepository(mongo)
	userService := services.NewUserService(userRepo)

	accountUseCase := usercase.AccountUsercase{
		AccountRepository:  repo,
		TransactionService: services.NewTransactionService(repo),
	}
	accountHandler := GinAccountHandlerImpl{accountUseCase}
	authHandler := NewAuthenticationHandlerGin(userService, tokenService)

	userUserCase := usercase.UserUsercase{AccountRepository: repo,
		UserService: userService}

	userHandler := NewGinUserhandler(userUserCase)

	handlerConfig := HandlerConfigGin{
		AccountHandler:        accountHandler,
		AuthenticationHandler: authHandler,
		UserHandler:           userHandler,
	}

	middlewareConfig := MiddlewareConfigGin{AuthenticationMiddleware: middleware.NewAuthenticationMiddlwareGin(tokenService)}
	serverConfig := ServerConfigurationGin{
		Port:                ":" + portNumber,
		HandlerConfig:       handlerConfig,
		MiddlewareConfigGin: middlewareConfig,
	}

	return serverConfig
}

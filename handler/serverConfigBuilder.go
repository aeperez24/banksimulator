package handler

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/middleware"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/services"
	"aeperez24/banksimulator/usercase"
)

func BuildServerConfig(portNumber string, tokenKey string, mongo config.MongoCofig) ServerConfiguration {
	repo := model.NewAccountMongoRepository(mongo)
	userRepo := model.NewUserMongoRepository(mongo)

	accountUseCase := usercase.AccountUsercase{
		AccountRepository:  repo,
		TransactionService: services.NewTransactionService(repo),
	}

	tokenService := services.NewTokenService(tokenKey)
	userService := services.NewUserService(userRepo)
	userUserCase := usercase.UserUsercase{AccountRepository: repo,
		UserService: userService}

	accountHandler := NewAccountHandler(accountUseCase)

	userHandler := NewUserhandler(userUserCase)
	authMiddleware := middleware.NewAuthenticationMiddlware(tokenService)
	authHandler := NewAuthenticationHandler(userService, tokenService)
	handlerConfig := HandlerConfig{
		AccountHandler:        accountHandler,
		UserHandler:           userHandler,
		AuthenticationHandler: authHandler,
	}

	serverConfig := ServerConfiguration{
		Port:             ":" + portNumber,
		MiddleWareConfig: middleware.MiddlewareConfig{AuthenticationMiddleware: authMiddleware},
		HandlerConfig:    handlerConfig,
	}

	return serverConfig
}

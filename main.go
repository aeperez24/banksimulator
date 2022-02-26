package main

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/handler"
	"aeperez24/banksimulator/middleware"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/services"
)

var DBConfig config.MongoCofig

func init() {
	if (config.MongoCofig{} == DBConfig) {
		DBConfig = config.BuildDBConfig()
	}

}

func main() {
	repo := model.NewAccountMongoRepository(DBConfig)
	userRepo := model.NewUserMongoRepository(DBConfig)
	accountHandler := handler.NewAccountHandler(repo)

	tokenService := services.NewTokenService("prodKey")
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserhandler(repo, userService)
	authMiddleware := middleware.NewAuthenticationMiddlware(tokenService)
	authHandler := handler.NewAuthenticationHandler(userService, tokenService)
	handlerConfig := handler.HandlerConfig{
		AccountHandler:        accountHandler,
		UserHandler:           userHandler,
		AuthenticationHandler: authHandler,
	}

	serverConfig := handler.ServerConfiguration{
		AccountHandler:   accountHandler,
		Port:             ":8080",
		MiddleWareConfig: middleware.MiddlewareConfig{AuthenticationMiddleware: authMiddleware},
		HandlerConfig:    handlerConfig,
	}
	server := handler.NewServer(serverConfig)
	server.Start()
}

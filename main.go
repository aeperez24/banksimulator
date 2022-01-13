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
	accountHandler := handler.NewAccountHandler(repo)
	tokenService := services.NewTokenService("prodKey")
	authMiddleware := middleware.NewAuthenticationMiddlware(tokenService)
	serverConfig := handler.ServerConfiguration{
		AccountHandler:   accountHandler,
		Port:             ":8080",
		MiddleWareConfig: middleware.MiddlewareConfig{AuthenticationMiddleware: authMiddleware},
	}
	server := handler.NewServer(serverConfig)
	server.Start()
}

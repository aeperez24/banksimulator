package handler

import (
	"aeperez24/banksimulator/middleware"
	"aeperez24/banksimulator/port"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerImpl struct {
	ServerConfiguration
	HttpServer http.Server
}

type ServerConfiguration struct {
	MiddleWareConfig middleware.MiddlewareConfig
	Port             string
	HandlerConfig    HandlerConfig
}

func NewServer(config ServerConfiguration) port.Server {
	return ServerImpl{ServerConfiguration: config}
}

func (mserver ServerImpl) Start() error {

	authMiddleware := mserver.MiddleWareConfig.AuthenticationMiddleware.Filter
	muxHandler := mux.NewRouter()
	accountHandler := mserver.HandlerConfig.AccountHandler
	userHandler := mserver.HandlerConfig.UserHandler
	authHandler := mserver.HandlerConfig.AuthenticationHandler
	muxHandler.HandleFunc("/account/balance/{AccountNumber}", authMiddleware(accountHandler.GetBalance))
	muxHandler.HandleFunc("/account/transfer/", authMiddleware(accountHandler.TransferMoney))
	muxHandler.HandleFunc("/account/deposit/", authMiddleware(accountHandler.Deposit))
	muxHandler.HandleFunc("/transaction/{AccountNumber}", authMiddleware(accountHandler.GetTransactions))
	muxHandler.HandleFunc("/user/signup", userHandler.CreateUser)
	muxHandler.HandleFunc("/user/signin", authHandler.Authenticate)

	mserver.HttpServer = http.Server{Addr: mserver.Port, Handler: muxHandler}
	return mserver.HttpServer.ListenAndServe()

}

func (mserver ServerImpl) Stop() {
	mserver.HttpServer.Shutdown(context.Background())
}

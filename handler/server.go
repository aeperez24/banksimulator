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
	AccountHandler   port.AccountHandler
	MiddleWareConfig middleware.MiddlewareConfig
	Port             string
}

func NewServer(config ServerConfiguration) port.Server {
	return ServerImpl{ServerConfiguration: config}
}

func (mserver ServerImpl) Start() {

	authMiddleware := mserver.MiddleWareConfig.AuthenticationMiddleware.Filter
	muxHandler := mux.NewRouter()

	muxHandler.HandleFunc("/balance/{AccountNumber}", authMiddleware(mserver.AccountHandler.GetBalance))
	muxHandler.HandleFunc("/transfer/", authMiddleware(mserver.AccountHandler.TransferMoney))
	muxHandler.HandleFunc("/deposit/", authMiddleware(mserver.AccountHandler.Deposit))
	muxHandler.HandleFunc("/transaction/{AccountNumber}", authMiddleware(mserver.AccountHandler.GetTransactions))

	mserver.HttpServer = http.Server{Addr: mserver.Port, Handler: muxHandler}
	err := mserver.HttpServer.ListenAndServe()
	if err != nil {
		println(err)
		panic(err)
	}

}

func (mserver ServerImpl) Stop() {
	mserver.HttpServer.Shutdown(context.Background())

}

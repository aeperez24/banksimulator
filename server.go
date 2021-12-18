package main

import (
	"aeperez24/banksimulator/port"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerImpl struct {
	AccountHandler port.AccountHandler
	Port           string
	HttpServer     http.Server
}

func NewServer(port string, accountHandler port.AccountHandler) port.Server {
	return ServerImpl{AccountHandler: accountHandler, Port: port}
}

func (mserver ServerImpl) Start() {
	muxHandler := mux.NewRouter()
	muxHandler.HandleFunc("/balance/{AccountNumber}", mserver.AccountHandler.GetBalance)
	muxHandler.HandleFunc("/transfer/", mserver.AccountHandler.TransferMoney)
	muxHandler.HandleFunc("/deposit/", mserver.AccountHandler.Deposit)
	mserver.HttpServer = http.Server{Addr: mserver.Port, Handler: muxHandler}
	mserver.HttpServer.ListenAndServe()

}
func (mserver ServerImpl) Stop() {
	mserver.HttpServer.Shutdown(context.Background())

}

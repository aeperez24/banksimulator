package main

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/handler"
)

var DBConfig config.MongoCofig

func init() {
	if (config.MongoCofig{} == DBConfig) {
		DBConfig = config.BuildDBConfig()
	}

}

func main() {
	serverConfig := handler.BuildServerConfigGin("8080", "prodKey", DBConfig)
	server := handler.NewGinServer(serverConfig)
	err := server.Start()
	if err != nil {
		println(err)
		panic(err)
	}
	server.Start()
}

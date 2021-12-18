package main

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/handler"
	"aeperez24/banksimulator/model"
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
	server := handler.NewServer(":8080", accountHandler)
	server.Start()
}

package main

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/handler"
	"log"
	"os"
)

var DBConfig config.MongoCofig

const ENVIRONMENTS_PATH = "envs/"
const LOCAL_ENVIRONMENT_CONFIG = "local"

func init() {
	env := os.Getenv("BANK_ENV")
	if env == "" {
		env = LOCAL_ENVIRONMENT_CONFIG
	}
	log.Println("running application with environment:" + env)
	config.LoadViperConfig(ENVIRONMENTS_PATH, env)

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

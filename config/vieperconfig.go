package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadViperConfig(path string, environment string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(environment)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	os.Setenv("DATABASE_HOST", viper.GetString("DATABASE_HOST"))
	os.Setenv("DATABASE_PORT", viper.GetString("DATABASE_PORT"))
	os.Setenv("DATABASE_USERNAME", viper.GetString("DATABASE_USERNAME"))
	os.Setenv("DATABASE_PASSWORD", viper.GetString("DATABASE_PASSWORD"))
	os.Setenv("JWT_SECRET", viper.GetString("JWT_SECRET"))

}

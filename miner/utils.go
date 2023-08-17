package Miner

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var config Config
	err = envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}

	return &config
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var config *Config

type Config struct {
	ServiceConfig  *ServiceConfig
	AuthConfig     *AuthConfig
	DBConfig       *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env file")
		os.Exit(1)
	}

	serviceConfig := ServiceConf()
	dbConfig := DBConf()
	authConfig := AuthConf()

	config = &Config{
		ServiceConfig:  serviceConfig,
		AuthConfig:     authConfig,
		DBConfig:       dbConfig,
	}

}

func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}
	return config
}

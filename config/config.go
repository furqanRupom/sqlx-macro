package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var config *Config

type AuthConfig struct {
	SecretToken string
}

type DBConfig struct {
	Host          string
	DBUser        string
	DBPass        string
	DBPort        int
	DBName        string
	SSLEnableMode bool
}

type Config struct {
	ServiceName string
	Version     string
	HttpPort    int
	AuthConfig  *AuthConfig
	DBConfig    *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env file")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Println("Service name is required")
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		log.Println("Version is required")
		os.Exit(1)
	}

	host := os.Getenv("HOST")
	if host == "" {
		log.Println("DB Host is required!")
		os.Exit(1)
	}
	dbUser := os.Getenv("DB_USER")

	if dbUser == "" {
		log.Println("DB User is required")
		os.Exit(1)
	}

	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		log.Println("DB Pass is required")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Println("DB name is required")
		os.Exit(1)
	}
	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		log.Println("DB port is required")
		os.Exit(1)
	}
	dbPrt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		log.Println("DB Port is should be a number")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Println("Server port is required!")
		os.Exit(1)
	}
	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		log.Println("Server Port is should be a number")
		os.Exit(1)
	}
	sslMode := os.Getenv("ENABLE_SSL_MODE")
	if sslMode == "" {
		log.Println("SSlMode is required")
		os.Exit(1)
	}
	sslModeBool, err := strconv.ParseBool(sslMode)
	if err != nil {
		log.Println("SSLMode must be a boolean value")
	}
	secret := os.Getenv("SECRET_TOKEN")
	if secret == "" {
		log.Println("Secret token is required!")
		os.Exit(1)
	}
	dbConfig := &DBConfig{
		DBUser:        dbUser,
		DBPass:        dbPass,
		Host:          host,
		DBPort:        int(dbPrt),
		DBName:        dbName,
		SSLEnableMode: sslModeBool,
	}

	authConfig := &AuthConfig{
		SecretToken: secret,
	}

	config = &Config{
		ServiceName: serviceName,
		Version:     version,
		HttpPort:    int(port),
		AuthConfig:  authConfig,
		DBConfig:    dbConfig,
	}

}

func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}
	return config
}

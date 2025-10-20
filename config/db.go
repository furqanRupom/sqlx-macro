package config

import (
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	Host          string
	DBUser        string
	DBPass        string
	DBPort        int
	DBName        string
	SSLEnableMode bool
}

func DBConf() *DBConfig {
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
	

	sslMode := os.Getenv("ENABLE_SSL_MODE")
	if sslMode == "" {
		log.Println("SSlMode is required")
		os.Exit(1)
	}
	sslModeBool, err := strconv.ParseBool(sslMode)
	if err != nil {
		log.Println("SSLMode must be a boolean value")
	}
	return &DBConfig{
		DBUser:        dbUser,
		DBPass:        dbPass,
		Host:          host,
		DBPort:        int(dbPrt),
		DBName:        dbName,
		SSLEnableMode: sslModeBool,
	}
}
package config

import (
	"log"
	"os"
	"strconv"
)

type ServiceConfig struct {
	ServiceName string
	Version     string
	HttpPort    int
}

func ServiceConf() *ServiceConfig {
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
	return &ServiceConfig{
		ServiceName: serviceName,
		Version:     version,
		HttpPort: int(port),
	}
}

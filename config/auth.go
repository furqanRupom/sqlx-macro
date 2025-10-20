package config

import (
	"log"
	"os"
)

type AuthConfig struct {
	SecretToken string
}

func AuthConf() *AuthConfig {
	secret := os.Getenv("SECRET_TOKEN")
	if secret == "" {
		log.Println("Secret token is required!")
		os.Exit(1)
	}
	return &AuthConfig{
		SecretToken: secret,
	}
}
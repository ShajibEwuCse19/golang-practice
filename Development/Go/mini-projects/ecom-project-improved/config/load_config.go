package config

import (
	"log"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

var configuration Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatal("HTTP_PORT is not set")
	}

	configuration = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
	}
}

func GetConfig() *Config {
	loadConfig()
	return &configuration
}

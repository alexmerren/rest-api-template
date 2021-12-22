package main

import (
	"golang-api-template/internal/config"
	"golang-api-template/internal/logger"
	"golang-api-template/internal/server"
)

func main() {
	err := initApp("config.json")
	if err != nil {
		panic(err)
	}
}

func initApp(configFileName string) error {
	config, err := config.ReadInConfig(configFileName)
	if err != nil {
		return err
	}

	logger, err := logger.NewZapLogger(config)
	if err != nil {
		return err
	}

	server, err := server.NewServer(config, logger)
	if err != nil {
		return err
	}

	server.Run()

	return nil
}

package main

import (
	"golang-api-template/internal/config"
	"golang-api-template/internal/logger"
	"golang-api-template/internal/server"
	"golang-api-template/internal/store"
)

func main() {
	err := initApp("config.json")
	if err != nil {
		panic(err)
	}
}

func initApp(configFileName string) error {
	err := config.ReadInConfig(configFileName)
	if err != nil {
		return err
	}

	logger, err := logger.NewZapLogger()
	if err != nil {
		return err
	}

	datastore, err := store.NewStore()
	if err != nil {
		return err
	}

	server, err := server.NewServer(logger, datastore)
	if err != nil {
		return err
	}

	server.Run()

	return nil
}

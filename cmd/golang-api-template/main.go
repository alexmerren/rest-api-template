package main

import (
	"golang-api-template/pkg/config"
	"golang-api-template/pkg/logger"
	"golang-api-template/pkg/server"
	"golang-api-template/pkg/store"
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

	store, err := store.NewStore()
	if err != nil {
		return err
	}

	server, err := server.NewServer(logger, store)
	if err != nil {
		return err
	}

	server.Run()

	return nil
}

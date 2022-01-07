package main

import (
	"golang-api-template/internal/config"
	"golang-api-template/internal/logger"
	"golang-api-template/internal/server"
	"golang-api-template/internal/store"
)

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}
}

func initApp() error {
	config, err := config.ReadInConfig()
	if err != nil {
		return err
	}

	logger, err := logger.NewZapLogger(config)
	if err != nil {
		return err
	}

	store, err := store.NewStore(config)
	if err != nil {
		return err
	}

	server, err := server.NewServer(config, logger, store)
	if err != nil {
		return err
	}

	server.Run()

	return nil
}

package main

import (
	"context"
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

	context := context.Background()

	store, err := store.NewStore(context, config)
	if err != nil {
		return err
	}
	defer store.CloseDB()

	server, err := server.NewServer(context, config, logger, store)
	if err != nil {
		return err
	}

	server.Run()

	return nil
}

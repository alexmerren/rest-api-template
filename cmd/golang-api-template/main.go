package main

import (
	"context"
	"golang-api-template/internal/config"
	"golang-api-template/internal/datastore"
	"golang-api-template/internal/logger"
	"golang-api-template/internal/router"
	"golang-api-template/internal/server"
)

func main() {
	app, err := ProvideApplication()
	if err != nil {
		panic(err)
	}
	app.Start()
	defer app.Stop()
}

type Application struct {
	logger    *logger.ZapLogger
	server    *server.Server
	datastore *datastore.Datastore
}

func ProvideApplication() (*Application, error) {
	filesystem := config.ProvideFilesystem()
	koanf := config.ProvideKoanf()
	config, err := config.ProvideConfig(koanf, filesystem)
	if err != nil {
		return nil, err
	}

	logger, err := logger.ProvideLogger(config)
	if err != nil {
		return nil, err
	}

	context := context.Background()
	datastore, err := datastore.ProvideDatastore(context, config)
	if err != nil {
		return nil, err
	}

	router := router.ProvideRouter(context, logger, datastore)
	server, err := server.ProvideServer(context, config, logger, datastore, router)
	if err != nil {
		return nil, err
	}

	return &Application{
		logger:    logger,
		server:    server,
		datastore: datastore,
	}, nil
}

func (a *Application) Start() {
	a.server.Run()
}

func (a *Application) Stop() {
	a.logger.Sync()
	_ = a.datastore.CloseDB()
}

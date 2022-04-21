package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"rest-api-template/internal/adapters/config"
	"rest-api-template/internal/adapters/logger"
	"rest-api-template/internal/adapters/memdb"
	"rest-api-template/internal/infrastructure/rest"
	"rest-api-template/internal/usecases"
)

func main() {
	ctx := context.Background()

	filesystem := config.NewFilesystem()
	config := config.NewConfiguration("config.yaml", filesystem)

	logLevel, _ := config.GetString("logger.loglevel")
	logger, err := logger.NewZapLogger(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Cleanup()

	memDB := memdb.NewMemoryStoreAdapter()
	usecases := usecases.NewRealContactUseCases(memDB, logger)
	port, _ := config.GetInt("server.port")
	server := rest.NewRESTServer(usecases, logger, port)

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt)

	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-cancelChan

	if err := server.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}

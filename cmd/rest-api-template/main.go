package main

import (
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
	filesystem := config.NewFilesystem()
	config := config.NewConfiguration("config.yaml", filesystem)
	logLevel, _ := config.GetString("LogLevel")
	logger, err := logger.NewZapLogger(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Cleanup()

	memDB := memdb.NewMemoryStoreAdapter()
	usecases := usecases.NewRealContactUseCases(memDB, logger)
	server := rest.NewRESTServer(usecases, logger)

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt)

	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-cancelChan
	logger.Debug("Program exiting...")
}

package main

import (
	"log"
	"os"
	"os/signal"
	"rest-api-template/internal/adapters/config"
	"rest-api-template/internal/infrastructure/database"
	"rest-api-template/internal/infrastructure/logger"
	"rest-api-template/internal/infrastructure/rest"
	"rest-api-template/internal/usecases"
)

func main() {
	filesystem := config.NewFilesystem()
	config := config.NewConfiguration("config.yaml", filesystem)

	logLevel, _ := config.GetString("logger.loglevel")
	logger, err := logger.NewZapLogger(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Cleanup()

	memDB := database.NewMemoryDatabase()
	usecases := usecases.NewRealContactUseCases(memDB, logger)

	port, _ := config.GetInt("server.port")
	server := rest.NewRESTServer(usecases, logger, port)

	terminationChannel := make(chan os.Signal, 1)
	signal.Notify(terminationChannel, os.Interrupt)

	if err = server.Start(); err != nil {
		logger.Error("failed to start HTTP server")
		os.Exit(1)
	}

	<-terminationChannel

	if err := server.Stop(); err != nil {
		logger.Error("error shutting down HTTP server")
		os.Exit(1)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"rest-api-template/internal/adapters/config"
	"rest-api-template/internal/adapters/logger"
	"rest-api-template/internal/adapters/memdb"
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/usecases"
)

func main() {
	ctx := context.Background()
	filesystem := config.NewFilesystem()
	config := config.NewConfiguration("config.yaml", filesystem)

	logLevel, err := config.GetString("LogLevel")
	if err != nil {
		log.Fatal(err)
	}
	logger, err := logger.NewZapLogger(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Cleanup()

	memDB := memdb.NewMemoryStoreAdapter()
	usecases := usecases.NewRealContactUseCases(memDB, logger)

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt)

	go func() {
		//err := server.Run()
		//if err != nil {
		//log.Fatalf(err)
		//}

		//TEST
		contact1, _ := entities.MakeContact("Alex", 22, "06/04/2000", "744 Filton Avenue", "Male")
		contact2, _ := entities.MakeContact("Ellie", 22, "24/12/1999", "Holly Cottage, Nomansland", "Female")
		newContacts := []*entities.Contact{contact1, contact2}
		usecases.CreateContacts(ctx, newContacts)

		contacts, _ := usecases.ListContacts(ctx)
		for _, contact := range contacts {
			fmt.Println(contact)
		}

		usecases.UpdateContactByID(ctx, contact1.ID, &entities.Contact{Address: "hahaha lmaooo homeless"})

		contacts, _ = usecases.ListContacts(ctx)
		for _, contact := range contacts {
			fmt.Println(contact)
		}
	}()

	<-cancelChan
	logger.Debug("Program exiting...")
}

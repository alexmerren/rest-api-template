package main

import (
	"context"
	"fmt"
	"rest-api-template/internal/adapters/config"
	"rest-api-template/internal/adapters/logger"
	"rest-api-template/internal/adapters/memdb"
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/usecases"
)

func main() {
	if err := NewService(); err != nil {
		panic(err)
	}
}

func NewService() error {
	ctx := context.Background()

	config := config.NewConfiguration()

	logLevel, err := config.GetString("LogLevel")
	if err != nil {
		return err
	}

	logger, err := logger.NewZapLogger(logLevel)
	if err != nil {
		return err
	}

	memDB := memdb.NewMemoryStoreAdapter()
	usecases := usecases.NewRealContactUseCases(memDB, logger)

	// THIS IS A TEST
	contact1, _ := entities.MakeContact("Alex", 22, "06/04/2000", "744 Filton Avenue", "Male")
	contact2, _ := entities.MakeContact("Ellie", 22, "24/12/1999", "Nomansland", "Female")
	newContacts := []entities.Contact{contact1, contact2}

	if err = usecases.CreateContacts(ctx, newContacts); err != nil {
		return err
	}

	returnedContacts, err := usecases.ListContacts(ctx)
	if err != nil {
		return err
	}

	fmt.Println(returnedContacts)

	return nil
}

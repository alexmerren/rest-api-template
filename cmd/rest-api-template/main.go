package main

import (
	"context"
	"rest-api-template/internal/adapters/memdb"
	"rest-api-template/internal/usecases"
)

func main() {
	ctx := context.Background()

	adapters := memdb.NewMemoryAdapter()
	usecases := usecases.NewRealContactUseCases(adapters)

	usecases.GetContactByID(ctx, "test")
}

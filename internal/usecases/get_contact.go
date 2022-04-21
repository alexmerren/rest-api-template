package usecases

import (
	"context"
	"fmt"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) GetContactByID(ctx context.Context, ID string) (entities.Contact, error) {
	fmt.Println("hi!")
	return entities.Contact{}, nil
}

package usecases

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) UpdateContactByID(ctx context.Context, ID string) (entities.Contact, error) {
	return entities.Contact{}, nil
}

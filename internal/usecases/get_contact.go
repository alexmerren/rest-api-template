package usecases

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) GetContactByID(ctx context.Context, ID string) (entities.Contact, error) {
	return u.store.ReadOne(ctx, ID)
}

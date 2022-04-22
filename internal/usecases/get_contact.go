package usecases

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) GetContactByID(ctx context.Context, ID string) (*entities.Contact, error) {
	contact, err := u.store.ReadContactWithID(ctx, ID)
	if err != nil {
		u.logger.Error(err)
		return nil, err
	}

	return contact, nil
}

package usecases

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) UpdateContactByID(ctx context.Context, ID string, newContact *entities.Contact) (*entities.Contact, error) {
	contact, err := u.store.UpdateContactWithID(ctx, ID, newContact)
	if err != nil {
		u.logger.Error(err)
		return nil, entities.NewInternalError("could not update Contact with ID", err)
	}

	return contact, nil
}

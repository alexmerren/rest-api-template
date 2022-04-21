package usecases

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) DeleteContactByID(ctx context.Context, ID string) (*entities.Contact, error) {
	contact, err := u.store.ReadOne(ctx, ID)
	if err != nil {
		u.logger.Error(err)
		return nil, entities.NewNotFoundError("could not find Contact with ID", err)
	}

	err = u.store.Delete(ctx, ID)
	if err != nil {
		u.logger.Error(err)
		return nil, entities.NewInternalError("could not delete Contact with ID", err)
	}

	return contact, nil
}

package usecases

import (
	"context"
	"fmt"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) DeleteContactByID(ctx context.Context, ID string) (*entities.Contact, error) {
	contact, err := u.store.ReadContactWithID(ctx, ID)
	if err != nil {
		u.logger.Error(err)
		return nil, entities.NewNotFoundError(fmt.Sprintf("could not find Contact with ID %s", ID), err)
	}

	err = u.store.DeleteContactWithID(ctx, ID)
	if err != nil {
		u.logger.Error(err)
		return nil, entities.NewNotFoundError(fmt.Sprintf("could not delete Contact with ID %s", ID), err)
	}

	return contact, nil
}

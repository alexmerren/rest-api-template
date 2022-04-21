package usecases

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) ListContacts(ctx context.Context) ([]*entities.Contact, error) {
	contacts, err := u.store.ReadMany(ctx)
	if err != nil {
		u.logger.Error(err)
		return nil, entities.NewInternalError("could not find contacts", err)
	}

	return contacts, nil
}

package usecases

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) CreateContacts(ctx context.Context, contacts []*entities.Contact) error {
	for _, contact := range contacts {
		if err := contact.Validate(); err != nil {
			u.logger.Error(err)
			return entities.NewBadRequestError("contact given was invalid", err)
		}

		err := u.store.CreateContact(ctx, contact)
		if err != nil {
			u.logger.Error(err)
			return entities.NewInternalError("could not create Contact", err)
		}
	}
	return nil
}

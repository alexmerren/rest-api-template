package usecases

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) CreateContacts(ctx context.Context, contacts []entities.Contact) error {
	for _, contact := range contacts {
		if err := contact.Validate(); err != nil {
			u.logger.Error(err)
			return err
		}

		err := u.store.Create(ctx, contact)
		if err != nil {
			u.logger.Error(err)
			return err
		}
	}
	return nil
}

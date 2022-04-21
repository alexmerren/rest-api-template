package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) ReadOne(ctx context.Context, ID string) (entities.Contact, error) {
	for _, contact := range m.contacts {
		if contact.ID == ID {
			return contact, nil
		}
	}

	return entities.Contact{}, entities.NewNotFoundError("a contact with that ID could not be found", nil)
}

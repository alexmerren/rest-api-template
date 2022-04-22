package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) ReadContactWithID(ctx context.Context, ID string) (*entities.Contact, error) {
	for _, contact := range m.Contacts {
		if contact.ID == ID {
			return contact, nil
		}
	}

	return nil, entities.NewNotFoundError("a contact with that ID could not be found", nil)
}

package memdb

import (
	"context"
	"fmt"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) ReadContactWithID(ctx context.Context, ID string) (*entities.Contact, error) {
	for _, contact := range m.Contacts {
		if contact.ID == ID {
			return contact, nil
		}
	}

	return nil, entities.NewNotFoundError(fmt.Sprintf("could not find contact with ID %s", ID), nil)
}

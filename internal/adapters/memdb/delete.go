package memdb

import (
	"context"
	"fmt"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) DeleteContactWithID(ctx context.Context, ID string) error {
	for index, contact := range m.Contacts {
		if contact.ID == ID {
			m.Contacts[index], m.Contacts[len(m.Contacts)-1] = m.Contacts[len(m.Contacts)-1], m.Contacts[index]
			m.Contacts = m.Contacts[:len(m.Contacts)-1]
			return nil
		}
	}

	return entities.NewNotFoundError(fmt.Sprintf("could not find contact with ID %s", ID), nil)
}

package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) Delete(ctx context.Context, ID string) error {
	for index, contact := range m.contacts {
		if contact.ID == ID {
			m.contacts[index], m.contacts[len(m.contacts)-1] = m.contacts[len(m.contacts)-1], m.contacts[index]
			m.contacts = m.contacts[:len(m.contacts)-1]
			return nil
		}
	}

	return entities.NewNotFoundError("could not find contact with given ID", nil)
}

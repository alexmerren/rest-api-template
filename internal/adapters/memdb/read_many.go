package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) ReadContacts(ctx context.Context) ([]*entities.Contact, error) {
	return m.Contacts, nil
}

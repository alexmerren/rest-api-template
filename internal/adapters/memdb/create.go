package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) Create(ctx context.Context, contact *entities.Contact) error {
	m.contacts = append(m.contacts, contact)
	return nil
}

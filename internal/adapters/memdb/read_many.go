package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) ReadMany(ctx context.Context) ([]entities.Contact, error) {
	return m.contacts, nil
}

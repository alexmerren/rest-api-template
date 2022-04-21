package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) Update(ctx context.Context, ID string, contact entities.Contact) (entities.Contact, error) {
	return entities.Contact{}, nil
}

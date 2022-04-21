package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryContactRepository) ReadOne(ctx context.Context, ID string) (entities.Contact, error) {
	return entities.Contact{}, nil
}

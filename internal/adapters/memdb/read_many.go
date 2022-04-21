package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryContactRepository) ReadMany(ctx context.Context) ([]entities.Contact, error) {
	return nil, nil
}

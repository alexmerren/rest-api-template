package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryContactRepository) Create(ctx context.Context, contact entities.Contact) error {
	return nil
}

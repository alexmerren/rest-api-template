package usecases

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (u *realContactUseCases) ListContacts(ctx context.Context) ([]entities.Contact, error) {
	return u.store.ReadMany(ctx)
}

package repositories

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

type Contacter interface {
	CreateContact(ctx context.Context, contact *entities.Contact) error
	DeleteContact(ctx context.Context, ID string) error
	UpdateContact(ctx context.Context, ID string, contact *entities.Contact) (*entities.Contact, error)
	GetContact(ctx context.Context, ID string) (*entities.Contact, error)
	ListContacts(ctx context.Context) ([]*entitites.Contact, error)
}

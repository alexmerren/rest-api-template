package repositories

import (
	"context"
	"net/http"
	"rest-api-template/internal/domain/entities"
)

type ContactStoreRepository interface {
	Create(ctx context.Context, contact *entities.Contact) error
	ReadOne(ctx context.Context, ID string) (*entities.Contact, error)
	ReadMany(ctx context.Context) ([]*entities.Contact, error)
	Update(ctx context.Context, ID string, contact *entities.Contact) (*entities.Contact, error)
	Delete(ctx context.Context, ID string) error
}

type ContactUseCases interface {
	CreateContacts(ctx context.Context, contacts []*entities.Contact) error
	GetContactByID(ctx context.Context, ID string) (*entities.Contact, error)
	ListContacts(ctx context.Context) ([]*entities.Contact, error)
	UpdateContactByID(ctx context.Context, ID string, newContact *entities.Contact) (*entities.Contact, error)
	DeleteContactByID(ctx context.Context, ID string) (*entities.Contact, error)
}

type ContactInfrastructure interface {
	Health(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	ReadOne(w http.ResponseWriter, r *http.Request)
	ReadMany(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

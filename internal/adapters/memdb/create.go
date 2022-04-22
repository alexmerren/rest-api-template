package memdb

import (
	"context"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) CreateContact(ctx context.Context, contact *entities.Contact) error {
	contact, err := entities.MakeContact(contact.Name, contact.Age, contact.Birthday, contact.Address, contact.Gender)
	if err != nil {
		return entities.NewInternalError("could not generate ID for Contact", err)
	}

	m.Contacts = append(m.Contacts, contact)
	return nil
}

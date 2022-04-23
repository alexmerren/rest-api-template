package memdb

import (
	"context"
	"fmt"
	"rest-api-template/internal/domain/entities"
)

func (m *memoryStoreAdapter) UpdateContactWithID(ctx context.Context, ID string, newContact *entities.Contact) (*entities.Contact, error) {
	// TODO this is a trash implementation, fix me!!
	for _, contact := range m.Contacts {
		if contact.ID == ID {
			if contact.Age != newContact.Age && newContact.Age != 0 {
				contact.Age = newContact.Age
			}

			if contact.Name != newContact.Name && newContact.Name != "" {
				contact.Name = newContact.Name
			}

			if contact.Birthday != newContact.Birthday && newContact.Birthday != "" {
				contact.Birthday = newContact.Birthday
			}

			if contact.Address != newContact.Address && newContact.Address != "" {
				contact.Address = newContact.Address
			}

			if contact.Gender != newContact.Gender && newContact.Gender != "" {
				contact.Gender = newContact.Gender
			}

			return contact, nil
		}
	}

	return nil, entities.NewNotFoundError(fmt.Sprintf("could not find contact with ID %s", ID), nil)
}

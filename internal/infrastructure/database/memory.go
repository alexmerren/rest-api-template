package database

import (
	"context"
	"fmt"
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/domain/repositories"
)

type memoryDatabase struct {
	Contacts []*entities.Contact
}

func NewMemoryDatabase() repositories.ContactStoreRepository {
	return &memoryDatabase{
		Contacts: make([]*entities.Contact, 0),
	}
}

func (m *memoryDatabase) CreateContact(ctx context.Context, contact *entities.Contact) error {
	contact, err := entities.NewContact(contact.Name, contact.Age, contact.Birthday, contact.Address, contact.Gender)
	if err != nil {
		return entities.NewInternalError("could not generate ID for Contact", err)
	}

	m.Contacts = append(m.Contacts, contact)
	return nil
}

func (m *memoryDatabase) DeleteContactWithID(ctx context.Context, ID string) error {
	for index, contact := range m.Contacts {
		if contact.ID == ID {
			m.Contacts[index], m.Contacts[len(m.Contacts)-1] = m.Contacts[len(m.Contacts)-1], m.Contacts[index]
			m.Contacts = m.Contacts[:len(m.Contacts)-1]
			return nil
		}
	}

	return entities.NewNotFoundError(fmt.Sprintf("could not find contact with ID %s", ID), nil)
}

func (m *memoryDatabase) ReadContacts(ctx context.Context) ([]*entities.Contact, error) {
	return m.Contacts, nil
}

func (m *memoryDatabase) ReadContactWithID(ctx context.Context, ID string) (*entities.Contact, error) {
	for _, contact := range m.Contacts {
		if contact.ID == ID {
			return contact, nil
		}
	}

	return nil, entities.NewNotFoundError(fmt.Sprintf("could not find contact with ID %s", ID), nil)
}

func (m *memoryDatabase) UpdateContactWithID(ctx context.Context, ID string, newContact *entities.Contact) (*entities.Contact, error) {
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

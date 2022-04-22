package memdb

import (
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/domain/repositories"
)

type memoryStoreAdapter struct {
	Contacts []*entities.Contact
}

func NewMemoryStoreAdapter() repositories.ContactStoreRepository {
	return &memoryStoreAdapter{
		Contacts: make([]*entities.Contact, 0),
	}
}

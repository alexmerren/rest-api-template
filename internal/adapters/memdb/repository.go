package memdb

import (
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/domain/repositories"
)

type memoryStoreAdapter struct {
	contacts []entities.Contact
}

func NewMemoryStoreAdapter() repositories.ContactStoreRepository {
	return &memoryStoreAdapter{
		contacts: make([]entities.Contact, 0),
	}
}

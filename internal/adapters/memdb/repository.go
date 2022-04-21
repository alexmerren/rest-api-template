package memdb

import (
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/domain/repositories"
)

type memoryContactRepository struct {
	contacts []entities.Contact
}

func NewMemoryAdapter() repositories.ContactRepository {
	return &memoryContactRepository{
		contacts: make([]entities.Contact, 0),
	}
}

package usecases

import "rest-api-template/internal/domain/repositories"

type realContactUseCases struct {
	contactRepo repositories.ContactRepository
}

func NewRealContactUseCases(contactRepository repositories.ContactRepository) repositories.ContactUseCases {
	return &realContactUseCases{
		contactRepo: contactRepository,
	}
}

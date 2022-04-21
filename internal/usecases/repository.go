package usecases

import "rest-api-template/internal/domain/repositories"

type realContactUseCases struct {
	store  repositories.ContactStoreRepository
	logger repositories.Logger
}

func NewRealContactUseCases(
	contactRepo repositories.ContactStoreRepository,
	loggerRepo repositories.Logger,
) repositories.ContactUseCases {
	return &realContactUseCases{
		store:  contactRepo,
		logger: loggerRepo,
	}
}

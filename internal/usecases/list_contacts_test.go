package usecases_test

import (
	"context"
	"errors"
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/infrastructure/logger"
	"rest-api-template/internal/usecases"
	"rest-api-template/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListContacts_HappyPath(t *testing.T) {
	//arrange
	ctx := context.Background()
	expectedContacts := []*entities.Contact{
		{
			ID:       testID,
			Name:     "Someone",
			Age:      22,
			Address:  "Somewhere",
			Gender:   "Something",
			Birthday: "Someday",
		},
	}
	mockLogger := new(mocks.Logger)
	mockAdapter := new(mocks.ContactStoreRepository)
	mockAdapter.On("ReadContacts", ctx).Return(expectedContacts, nil).Once()
	defer mockAdapter.AssertExpectations(t)
	usecases := usecases.NewRealContactUseCases(mockAdapter, mockLogger)

	//act
	actualContacts, err := usecases.ListContacts(ctx)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, expectedContacts, actualContacts)
}

func TestListContacts_AdapterError(t *testing.T) {
	//arrange
	ctx := context.Background()
	logger, err := logger.NewZapLogger("debug")
	assert.NoError(t, err)
	expectedError := errors.New("mock error")
	mockAdapter := new(mocks.ContactStoreRepository)
	mockAdapter.On("ReadContacts", ctx).Return(nil, expectedError).Once()
	defer mockAdapter.AssertExpectations(t)
	usecases := usecases.NewRealContactUseCases(mockAdapter, logger)

	//act
	contacts, err := usecases.ListContacts(ctx)

	//assert
	assert.Error(t, err)
	assert.Nil(t, contacts)
	assert.IsType(t, &entities.InternalError{}, err)
}

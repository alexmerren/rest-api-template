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

const testAddress = "test-new-address-location"

func TestUpdateContactByID_HappyPath(t *testing.T) {
	//arrange
	ctx := context.Background()
	expectedContact := &entities.Contact{
		Name:     "Someone",
		Age:      22,
		Address:  testAddress,
		Gender:   "Something",
		Birthday: "Someday",
	}
	partialContact := &entities.Contact{
		Address: testAddress,
	}
	mockLogger := new(mocks.Logger)
	mockAdapter := new(mocks.ContactStoreRepository)
	mockAdapter.On("UpdateContactWithID", ctx, testID, partialContact).Return(expectedContact, nil).Once()
	defer mockAdapter.AssertExpectations(t)
	usecases := usecases.NewRealContactUseCases(mockAdapter, mockLogger)

	//act
	returnedContact, err := usecases.UpdateContactByID(ctx, testID, partialContact)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, expectedContact, returnedContact)
}

func TestUpdateContactByID_AdapterError(t *testing.T) {
	//arrange
	ctx := context.Background()
	testErr := errors.New("test error")
	logger, err := logger.NewZapLogger("debug")
	assert.NoError(t, err)
	partialContact := &entities.Contact{
		Address: testAddress,
	}
	mockAdapter := new(mocks.ContactStoreRepository)
	mockAdapter.On("UpdateContactWithID", ctx, testID, partialContact).Return(nil, testErr).Once()
	defer mockAdapter.AssertExpectations(t)
	usecases := usecases.NewRealContactUseCases(mockAdapter, logger)

	//act
	contact, err := usecases.UpdateContactByID(ctx, testID, partialContact)

	//assert
	assert.NotNil(t, err)
	assert.Nil(t, contact)
}

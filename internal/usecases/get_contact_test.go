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

const testID = "test-contact-id"

func TestGetContactByID_HappyPath(t *testing.T) {
	//arrange
	ctx := context.Background()
	contact := &entities.Contact{
		Name:     "Someone",
		Age:      22,
		Address:  "Somewhere",
		Gender:   "Something",
		Birthday: "Someday",
	}
	mockLogger := new(mocks.Logger)
	mockAdapter := new(mocks.ContactStoreRepository)
	mockAdapter.On("ReadContactWithID", ctx, testID).Return(contact, nil).Once()
	defer mockAdapter.AssertExpectations(t)
	usecases := usecases.NewRealContactUseCases(mockAdapter, mockLogger)

	//act
	returnedContact, err := usecases.GetContactByID(ctx, testID)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, contact, returnedContact)
}

func TestGetContactByID_AdapterErr(t *testing.T) {
	//arrange
	ctx := context.Background()
	testErr := errors.New("test error")
	logger, err := logger.NewZapLogger("debug")
	assert.NoError(t, err)

	mockAdapter := new(mocks.ContactStoreRepository)
	mockAdapter.On("ReadContactWithID", ctx, testID).Return(nil, testErr).Once()
	defer mockAdapter.AssertExpectations(t)
	usecases := usecases.NewRealContactUseCases(mockAdapter, logger)

	//act
	contact, err := usecases.GetContactByID(ctx, testID)

	//assert
	assert.NotNil(t, err)
	assert.Nil(t, contact)
}

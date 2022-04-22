package usecases_test

import (
	"context"
	"errors"
	"rest-api-template/internal/adapters/logger"
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/usecases"
	"rest-api-template/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateContacts_HappyPath(t *testing.T) {
	//arrange
	ctx := context.Background()
	contacts := []*entities.Contact{
		{
			Name:     "Someone",
			Age:      22,
			Address:  "Somewhere",
			Gender:   "Something",
			Birthday: "Someday",
		},
	}
	mockLogger := new(mocks.Logger)
	mockAdapter := new(mocks.ContactStoreRepository)
	mockAdapter.On("Create", ctx, contacts[0]).Return(nil).Once()
	defer mockAdapter.AssertExpectations(t)
	usecases := usecases.NewRealContactUseCases(mockAdapter, mockLogger)

	//act
	err := usecases.CreateContacts(ctx, contacts)

	//assert
	assert.NoError(t, err)
}

func TestCreateContacts_AdapterError(t *testing.T) {
	//arrange
	ctx := context.Background()
	logger, err := logger.NewZapLogger("debug")
	assert.NoError(t, err)
	testErr := errors.New("mock error")
	contacts := []*entities.Contact{
		{
			Name:     "Someone",
			Age:      22,
			Address:  "Somewhere",
			Gender:   "Something",
			Birthday: "Someday",
		},
	}
	mockAdapter := new(mocks.ContactStoreRepository)
	mockAdapter.On("Create", ctx, contacts[0]).Return(testErr).Once()
	defer mockAdapter.AssertExpectations(t)
	usecases := usecases.NewRealContactUseCases(mockAdapter, logger)

	//act
	err = usecases.CreateContacts(ctx, contacts)

	//assert
	assert.NotNil(t, err)
	assert.IsType(t, &entities.InternalError{}, err)
}

func TestCreateContacts_ValidationError(t *testing.T) {
	//arrange
	ctx := context.Background()
	logger, err := logger.NewZapLogger("debug")
	assert.NoError(t, err)
	invalidContacts := []*entities.Contact{
		{
			Name:     "Someone",
			Age:      22,
			Gender:   "Something",
			Birthday: "Someday",
		},
	}
	usecases := usecases.NewRealContactUseCases(nil /*adapter*/, logger)

	//act
	err = usecases.CreateContacts(ctx, invalidContacts)

	//assert
	assert.NotNil(t, err)
	assert.IsType(t, &entities.BadRequestError{}, err)
	assert.Equal(t, "contact given was invalid", err.Error())
}

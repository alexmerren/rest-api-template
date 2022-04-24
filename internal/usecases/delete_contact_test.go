package usecases_test

import (
	"context"
	"errors"
	"fmt"
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/infrastructure/logger"
	"rest-api-template/internal/usecases"
	"rest-api-template/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteContactByID_HappyPath(t *testing.T) {
	//arrange
	ctx := context.Background()
	expectedContact := &entities.Contact{
		Name:     "Someone",
		Age:      22,
		Address:  "Somewhere",
		Gender:   "Something",
		Birthday: "Someday",
	}
	mockLogger := new(mocks.Logger)
	mockAdapter := new(mocks.ContactStoreRepository)
	mockAdapter.On("ReadContactWithID", ctx, testID).Return(expectedContact, nil).Once()
	mockAdapter.On("DeleteContactWithID", ctx, testID).Return(nil).Once()
	defer mockAdapter.AssertExpectations(t)
	usecases := usecases.NewRealContactUseCases(mockAdapter, mockLogger)

	//act
	contact, err := usecases.DeleteContactByID(ctx, testID)

	//assert
	assert.NoError(t, err)
	assert.NotNil(t, contact)
	assert.Equal(t, expectedContact, contact)
}

func TestDeleteContactByID_AdapterErrors(t *testing.T) {
	//arrange
	ctx := context.Background()
	logger, err := logger.NewZapLogger("debug")
	assert.NoError(t, err)
	mockAdapter := new(mocks.ContactStoreRepository)
	defer mockAdapter.AssertExpectations(t)
	usecases := usecases.NewRealContactUseCases(mockAdapter, logger)

	var testCases = []struct {
		name               string
		expectedErrMsg     string
		readContactError   error
		deleteContactError error
	}{
		{
			name:             "Read Conctact Error",
			readContactError: errors.New("mock error"),
			expectedErrMsg:   fmt.Sprintf("could not find Contact with ID %s", testID),
		},
		{
			name:               "Delete Conctact Error",
			deleteContactError: errors.New("mock error"),
			expectedErrMsg:     fmt.Sprintf("could not delete Contact with ID %s", testID),
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			if tc.readContactError != nil {
				mockAdapter.On("ReadContactWithID", ctx, testID).Return(nil, tc.readContactError).Once()
			} else {
				mockAdapter.On("ReadContactWithID", ctx, testID).Return(&entities.Contact{}, nil).Once()
			}

			if tc.deleteContactError != nil {
				mockAdapter.On("DeleteContactWithID", ctx, testID).Return(tc.deleteContactError).Once()
			}

			//act
			contact, err := usecases.DeleteContactByID(ctx, testID)

			//assert
			assert.NotNil(t, err)
			assert.Nil(t, contact)
			assert.Equal(t, tc.expectedErrMsg, err.Error())
		})
	}
}

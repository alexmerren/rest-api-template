package datastore_test

import (
	"context"
	"errors"
	"rest-api-template/internal/datastore"
	"rest-api-template/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	validUsername = "test"
	validPassword = "test"
	validPort     = 8000
	validName     = "test"
	validHost     = "test"
)

func Test_ProvideDatastore_HappyPath(t *testing.T) {}

func Test_ProvideDatastore_ConfigErrors(t *testing.T) {
	testCases := []struct {
		name           string
		userErr        error
		passErr        error
		portErr        error
		nameErr        error
		hostErr        error
		expectedErrMsg string
	}{
		{
			name:           "username config error",
			userErr:        errors.New("mock error"),
			expectedErrMsg: "mock error",
		},
		{
			name:           "password config error",
			passErr:        errors.New("mock error"),
			expectedErrMsg: "mock error",
		},
		{
			name:           "port config error",
			portErr:        errors.New("mock error"),
			expectedErrMsg: "mock error",
		},
		{
			name:           "name config error",
			nameErr:        errors.New("mock error"),
			expectedErrMsg: "mock error",
		},
		{
			name:           "host config error",
			hostErr:        errors.New("mock error"),
			expectedErrMsg: "mock error",
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			mockConfig := new(mocks.Config)

			if tc.userErr != nil {
				mockConfig.On("GetString", "Database.Username").Return("", tc.userErr)
			} else {
				mockConfig.On("GetString", "Database.Username").Return(validUsername, nil)
			}

			if tc.passErr != nil {
				mockConfig.On("GetString", "Database.Password").Return("", tc.passErr)
			} else {
				mockConfig.On("GetString", "Database.Password").Return(validPassword, nil)
			}

			if tc.portErr != nil {
				mockConfig.On("GetInt", "Database.Port").Return(0, tc.portErr)
			} else {
				mockConfig.On("GetInt", "Database.Port").Return(validPort, nil)
			}

			if tc.nameErr != nil {
				mockConfig.On("GetString", "Database.Name").Return("", tc.nameErr)
			} else {
				mockConfig.On("GetString", "Database.Name").Return(validName, nil)
			}

			if tc.hostErr != nil {
				mockConfig.On("GetString", "Host").Return("", tc.hostErr)
			} else {
				mockConfig.On("GetString", "Host").Return(validHost, nil)
			}

			store, storeErr := datastore.ProvideDatastore(context.Background(), mockConfig)
			assert.Nil(t, store)
			assert.EqualError(t, storeErr, "mock error")
		})
	}
}

func Test_ProvideDatastore_OpenError(t *testing.T) {}

func Test_ProvideDatastore_PingError(t *testing.T) {}

func Test_Datastore_QueryContext_HappyPath(t *testing.T) {}

func Test_Datastore_QueryContext_Error(t *testing.T) {}

func Test_Datastore_QueryRowContext_HappyPath(t *testing.T) {}

func Test_Datastore_QueryRowContext_Error(t *testing.T) {}

func Test_Datastore_ExecContext_HappyPath(t *testing.T) {}

func Test_Datastore_ExecContext_Errors(t *testing.T) {}

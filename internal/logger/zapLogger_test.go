package logger_test

import (
	"errors"
	"golang-api-template/internal/config"
	"golang-api-template/internal/logger"
	"golang-api-template/mocks"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const validConfigJSON = `
	{
		"logger": {
			"level" : "info", 
			"encoding" : "console"
		}
	}
`

func Test_ProvideLogger_HappyPath(t *testing.T) {
	testFilesystem := fstest.MapFS{
		"config.json": {
			Data: []byte(validConfigJSON),
			Mode: 0644,
		},
	}
	koanf := config.ProvideKoanf()

	configStruct, configErr := config.ProvideConfig(koanf, testFilesystem)
	require.NoError(t, configErr, "ProvideConfig should not return an error")
	zapLogger, loggerErr := logger.ProvideLogger(configStruct)

	require.NotNil(t, zapLogger, "ProvideLogger did not return logger unexpectedly")
	require.Implements(t, (*logger.Logger)(nil), zapLogger, "ProvideLogger returned wrong type")
	require.NoError(t, loggerErr, "ProvideLogger returned error unexpectedly")
}

func Test_ProvideLogger_ConfigErrors(t *testing.T) {
	testCases := []struct {
		name           string
		levelErr       error
		encodingErr    error
		expectedErrMsg string
	}{
		{
			name:           "level value error",
			levelErr:       errors.New("mock error"),
			expectedErrMsg: "mock error",
		},
		{
			name:           "encoding value error",
			encodingErr:    errors.New("mock error"),
			expectedErrMsg: "mock error",
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			mockConfig := new(mocks.Config)

			if tc.levelErr != nil {
				mockConfig.On("GetString", "logger.level").Return("", tc.levelErr)
			} else {
				mockConfig.On("GetString", "logger.level").Return("info", nil)
			}

			if tc.encodingErr != nil {
				mockConfig.On("GetString", "logger.encoding").Return("", tc.encodingErr)
			} else {
				mockConfig.On("GetString", "logger.encoding").Return("console", nil)
			}

			zapLogger, loggerErr := logger.ProvideLogger(mockConfig)

			assert.Nil(t, zapLogger, "")
			assert.EqualError(t, loggerErr, tc.expectedErrMsg, "")
		})
	}
}

func Test_ProvideLogger_BuildError(t *testing.T) {}

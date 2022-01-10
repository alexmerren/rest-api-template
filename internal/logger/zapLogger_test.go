package logger

import (
	"golang-api-template/internal/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewZapLogger_HappyPath(t *testing.T) {
	os.Setenv("logger_encoding", "json")
	os.Setenv("logger_level", "info")
	defer os.Clearenv()
	config, err := config.NewConfig("")
	require.NoError(t, err)

	logger, err := NewZapLogger(config)
	assert.NotNil(t, logger)
	assert.NoError(t, err)
}

func TestNewZapLogger_InvalidLogLevel(t *testing.T) {
	os.Setenv("logger_level", "INVALID_STRING")
	defer os.Clearenv()
	config, err := config.NewConfig("")
	require.NoError(t, err)

	logger, err := NewZapLogger(config)
	assert.Nil(t, logger)
	assert.Error(t, err)
	assert.Equal(t, "unrecognized level: \"INVALID_STRING\"", err.Error())
}

func TestNewZapLogger_InvalidEncoding(t *testing.T) {
	os.Setenv("logger_encoding", "INVALID_STRING")
	defer os.Clearenv()
	config, err := config.NewConfig("")
	require.NoError(t, err)

	logger, err := NewZapLogger(config)
	assert.Nil(t, logger)
	assert.Error(t, err)
	assert.Equal(t, "unrecognized level: \"INVALID_STRING\"", err.Error())
}

func TestNewZapLogger_ConfigNotFoundError(t *testing.T) {
	os.Setenv("logger_encoding", "")
	os.Setenv("logger_level", "")
	defer os.Clearenv()
	config, err := config.NewConfig("")
	require.NoError(t, err)

	logger, err := NewZapLogger(config)
	assert.Nil(t, logger)
	assert.Error(t, err)
	assert.Equal(t, "Could not find value logger.encoding", err.Error())
}

package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const validJson = `
{ "test": "hello!" }
`

func Test_NewConfig_HappyPath(t *testing.T) {
	testFilename := "test-config.json"
	cleanup := setupFileSystem(t, testFilename, validJson, false /* fileErr */)
	defer cleanup()

	config, err := NewConfig(testFilename)
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Implements(t, (*ConfigInterface)(nil), config)
}

func Test_NewConfig_NoFailureWithMissingFile(t *testing.T) {
	cleanup := setupFileSystem(t, "", "", true /* fileErr */)
	defer cleanup()

	config, err := NewConfig("file-does-not-exist.json")
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func Test_NewConfig_FailureReadingFile(t *testing.T) {
	testFilename := "test-config.json"
	cleanup := setupFileSystem(t, testFilename, "not json", false /* fileErr */)
	defer cleanup()

	config, err := NewConfig(testFilename)
	assert.Nil(t, config)
	assert.Contains(t, err.Error(), "error reading config file:")
}

func Test_GetString(t *testing.T) {}

func Test_GetInt(t *testing.T) {}

func setupFileSystem(t *testing.T, filename string, data string, fileErr bool) func() {
	cleanupFunc := func() {}
	if fileErr == false {
		ioutil.WriteFile(filename, []byte(data), 0644)
		cleanupFunc = func() {
			os.Remove(filename)
		}
	}
	return cleanupFunc
}

package config_test

import (
	"errors"
	"golang-api-template/internal/config"
	"golang-api-template/mocks"
	"io/fs"
	"os"
	"testing"
	"testing/fstest"

	"github.com/knadh/koanf/parsers/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const validJSON = `{ "port": 8080 }`

func Test_ProvideConfig_File_HappyPath(t *testing.T) {
	testFilesystem := fstest.MapFS{
		"config.json": {
			Data: []byte(validJSON),
			Mode: 0644,
		},
	}
	koanf := config.ProvideKoanf()

	configStruct, configErr := config.ProvideConfig(koanf, testFilesystem)
	portValue, portErr := configStruct.GetInt("port")

	assert.NoError(t, configErr, "ProvideConfig returned an error unexpectedly")
	assert.NotNil(t, configStruct, "ProvideConfig returned StructConfig unexpectedly")
	assert.IsType(t, configStruct, &config.MapConfig{}, "ProvideConfig returned wrong type")
	assert.NoError(t, portErr, "GetInt shoudl not return error")
	assert.Equal(t, 8080, portValue, "config.Port should be 8080")
}

func Test_ProvideConfig_File_ReadingError(t *testing.T) {
	mockFilesystem := new(mocks.Filesystem)
	mockFilesystem.On("Open", "config.json").Return(nil, nil)

	mockKoanf := new(mocks.Koanf)
	mockKoanf.On("Load", mock.AnythingOfType("*fs.FS"), json.Parser()).Return(errors.New("mock error"))
	mockKoanf.On("Load", mock.AnythingOfType("*env.Env"), nil).Return(nil).Run(func(args mock.Arguments) {})
	mockKoanf.On("Unmarshal", "", &config.MapConfig{}).Return(nil)

	configStruct, configErr := config.ProvideConfig(mockKoanf, mockFilesystem)

	assert.EqualError(t, configErr, "error reading config file: mock error", "ProvideConfig did not return an error unexpectedly")
	assert.Nil(t, configStruct, "ProvideConfig returned StructConfig unexpectedly")
}

func Test_ProvideConfig_Env_HappyPath(t *testing.T) {
	os.Setenv("LOGGER_ENCODING", "console")
	defer os.Clearenv()

	koanf := config.ProvideKoanf()
	filesystem := config.ProvideFilesystem()
	config, configErr := config.ProvideConfig(koanf, filesystem)
	encodingValue, encodingErr := config.GetString("logger.encoding")

	assert.NoError(t, configErr, "")
	assert.NotNil(t, config, "")
	assert.NoError(t, encodingErr, "GetString should not return error")
	assert.Equal(t, "console", encodingValue, "logger.encoding should be \"console\"")
}

func Test_ProvideConfig_Env_ReadingError(t *testing.T) {
	filesystem := config.ProvideFilesystem()
	mockKoanf := new(mocks.Koanf)
	mockKoanf.On("Load", mock.AnythingOfType("*fs.FS"), json.Parser()).Return(nil)
	mockKoanf.On("Load", mock.AnythingOfType("*env.Env"), nil).Return(errors.New("mock error")).Run(func(args mock.Arguments) {})
	mockKoanf.On("Unmarshal", "", &config.MapConfig{}).Return(nil)

	configStruct, configErr := config.ProvideConfig(mockKoanf, filesystem)

	assert.Nil(t, configStruct, "ProvideConfig returned StructConfig unexpectedly")
	assert.EqualError(t, configErr, "error reading environment variables: mock error", "ProvideConfig should return an error")
}

func Test_ProvideConfig_UnmarshalError(t *testing.T) {
	filesystem := config.ProvideFilesystem()
	mockKoanf := new(mocks.Koanf)
	mockKoanf.On("Load", mock.AnythingOfType("*fs.FS"), json.Parser()).Return(nil)
	mockKoanf.On("Load", mock.AnythingOfType("*env.Env"), nil).Return(nil).Run(func(args mock.Arguments) {})
	mockKoanf.On("Unmarshal", "", &config.MapConfig{}).Return(errors.New("mock error"))

	configStruct, configErr := config.ProvideConfig(mockKoanf, filesystem)

	assert.Nil(t, configStruct, "ProvideConfig returned StructConfig unexpectedly")
	assert.EqualError(t, configErr, "error unmarshalling to config struct: mock error", "ProvideConfig should return an error")
}

func Test_ProveKoanf(t *testing.T) {
	koanf := config.ProvideKoanf()
	assert.Implements(t, (*config.Koanf)(nil), koanf, "wrong type returned")
}

func Test_ProvideFileSystem(t *testing.T) {
	filesystem := config.ProvideFilesystem()
	assert.Implements(t, (*fs.FS)(nil), filesystem, "wrong type returned")
}

func Test_StructConfig_GetString_HappyPath(t *testing.T) {
	filesystem := config.ProvideFilesystem()
	mockKoanf := new(mocks.Koanf)
	mockKoanf.On("Load", mock.AnythingOfType("*fs.FS"), json.Parser()).Return(errors.New("mock error"))
	mockKoanf.On("Load", mock.AnythingOfType("*env.Env"), nil).Return(nil).Run(func(args mock.Arguments) {})
	mockKoanf.On("Unmarshal", "", &config.MapConfig{}).Return(nil)
	mockKoanf.On("String", "testvalue").Return("testValueReturn")

	configStruct, configErr := config.ProvideConfig(mockKoanf, filesystem)
	require.NoError(t, configErr)

	value, err := configStruct.GetString("testvalue")
	assert.Equal(t, "testValueReturn", value, "GetString did not return a value unexpectedly")
	assert.NoError(t, err, "could not find value testvalue", "GetString returned an error unexpectedly")
}

func Test_StructConfig_GetString_ValueNotFoundError(t *testing.T) {
	filesystem := config.ProvideFilesystem()
	mockKoanf := new(mocks.Koanf)
	mockKoanf.On("Load", mock.AnythingOfType("*fs.FS"), json.Parser()).Return(errors.New("mock error"))
	mockKoanf.On("Load", mock.AnythingOfType("*env.Env"), nil).Return(nil).Run(func(args mock.Arguments) {})
	mockKoanf.On("Unmarshal", "", &config.MapConfig{}).Return(nil)
	mockKoanf.On("String", "testvalue").Return("")

	configStruct, configErr := config.ProvideConfig(mockKoanf, filesystem)
	require.NoError(t, configErr)

	value, err := configStruct.GetString("testvalue")
	assert.Equal(t, value, "", "GetString returned a value unexpectedly")
	assert.EqualError(t, err, "could not find value testvalue", "GetString did not return an error unexpectedly")

}

func Test_StructConfig_GetInt_HappyPath(t *testing.T) {
	filesystem := config.ProvideFilesystem()
	mockKoanf := new(mocks.Koanf)
	mockKoanf.On("Load", mock.AnythingOfType("*fs.FS"), json.Parser()).Return(errors.New("mock error"))
	mockKoanf.On("Load", mock.AnythingOfType("*env.Env"), nil).Return(nil).Run(func(args mock.Arguments) {})
	mockKoanf.On("Unmarshal", "", &config.MapConfig{}).Return(nil)
	mockKoanf.On("Int", "testvalue").Return(1)

	configStruct, configErr := config.ProvideConfig(mockKoanf, filesystem)
	require.NoError(t, configErr)

	value, err := configStruct.GetInt("testvalue")
	assert.Equal(t, value, 1, "GetString did not return a value unexpectedly")
	assert.NoError(t, err, "could not find value testvalue", "GetString returned an error unexpectedly")
}

func Test_StructConfig_GetInt_ValueNotFoundError(t *testing.T) {
	filesystem := config.ProvideFilesystem()
	mockKoanf := new(mocks.Koanf)
	mockKoanf.On("Load", mock.AnythingOfType("*fs.FS"), json.Parser()).Return(errors.New("mock error"))
	mockKoanf.On("Load", mock.AnythingOfType("*env.Env"), nil).Return(nil).Run(func(args mock.Arguments) {})
	mockKoanf.On("Unmarshal", "", &config.MapConfig{}).Return(nil)
	mockKoanf.On("Int", "testvalue").Return(0)

	configStruct, configErr := config.ProvideConfig(mockKoanf, filesystem)
	require.NoError(t, configErr)

	value, err := configStruct.GetInt("testvalue")
	assert.Equal(t, value, 0, "GetInt returned a value unexpectedly")
	assert.EqualError(t, err, "could not find value testvalue", "GetInt did not return an error unexpectedly")

}

package config

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/env"
	fsys "github.com/knadh/koanf/providers/fs"
)

var k Koanf

type MapConfig struct{}

// This Koanf interface is used to create mocks for testing
type Koanf interface {
	Load(p koanf.Provider, pa koanf.Parser, opts ...koanf.Option) error
	Unmarshal(path string, o interface{}) error
	String(name string) string
	Int(name string) int
}

type Filesystem interface {
	Open(name string) (fs.File, error)
}

// NewConfig takes in a filename and unmarshals it into a config struct.
func ProvideConfig(koanf Koanf, filesystem Filesystem) (Config, error) {
	filename := "config.json"
	k = koanf

	// Check if the file exists, if not, just use the environment variables.
	// Load in the config file using the json parser to read the variables into the config struct
	if _, err := filesystem.Open(filename); err == nil {
		if err := k.Load(fsys.Provider(filesystem, filename), json.Parser()); err != nil {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Load in the environment variables that correspond to what we want in the config struct
	err := k.Load(env.Provider("" /* Prefix */, "." /* Delimeter */, func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil /* Parser */)
	if err != nil {
		return nil, fmt.Errorf("error reading environment variables: %w", err)
	}

	config := &MapConfig{}
	err = k.Unmarshal("" /* path */, config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling to config struct: %v", err)
	}

	return config, nil
}

func ProvideKoanf() Koanf {
	return koanf.New(".")
}

func ProvideFilesystem() Filesystem {
	return os.DirFS(".")
}

func (c *MapConfig) GetString(name string) (string, error) {
	value := k.String(strings.ToLower(name))
	if value == "" {
		return "", fmt.Errorf("could not find value %v", name)
	}
	return value, nil
}

func (c *MapConfig) GetInt(name string) (int, error) {
	value := k.Int(strings.ToLower(name))
	if value == 0 {
		return 0, fmt.Errorf("could not find value %v", name)
	}
	return value, nil
}

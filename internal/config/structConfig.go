package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

var k = koanf.New(".")

type StructConfig struct {
	configFile string
	Host       string         `koanf:"host"`
	Port       int            `koanf:"port"`
	Logger     LoggerConfig   `koanf:"logger"`
	Database   DatabaseConfig `koanf:"database"`
}

type LoggerConfig struct {
	Level    string `koanf:"level"`
	Encoding string `koanf:"encoding"`
}

type DatabaseConfig struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Name     string `koanf:"name"`
	Port     int    `koanf:"port"`
}

// NewConfig takes in a filename and unmarshals it into a config struct.
// nolint:errcheck // Environment provider can never return an error so we ignore it.
func NewConfig(filename string) (*StructConfig, error) {
	config := &StructConfig{
		configFile: filename,
	}

	if err := config.init(); err != nil {
		return nil, err
	}

	err := k.Unmarshal("" /* path */, config)
	if err != nil {
		return nil, err
	}

	spew.Dump(config)

	return config, nil
}

func (s *StructConfig) init() error {
	if _, err := os.Stat(s.configFile); err == nil {
		// Load in a config file with the given name using the json parser.
		if err := k.Load(file.Provider(s.configFile), json.Parser()); err != nil {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Load in the environment variables that correspond to what we want
	err := k.Load(env.Provider("" /* Prefix */, "." /* Delimeter */, func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil /* Parser */)
	if err != nil {
		return fmt.Errorf("error reading environment variables: %w", err)
	}

	return nil
}

func (s *StructConfig) GetString(name string) (string, error) {
	value := k.String(strings.ToLower(name))
	if value == "" {
		return "", fmt.Errorf("Could not find value %v", name)
	}
	return value, nil
}

func (s *StructConfig) GetInt(name string) (int, error) {
	value := k.Int(strings.ToLower(name))
	if value == 0 {
		return 0, fmt.Errorf("Could not find value %v", name)
	}
	return value, nil
}

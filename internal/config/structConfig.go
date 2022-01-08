package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

var k = koanf.New(".")

type StructConfig struct {
	Host     string         `koanf:"host"`
	Port     int            `koanf:"port"`
	Logger   LoggerConfig   `koanf:"logger"`
	Database DatabaseConfig `koanf:"database"`
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

// ReadInConfig takes in a filename and unmarshals it into a config struct.
// nolint:errcheck // Environment provider can never return an error so we ignore it.
func ReadInConfig() (*StructConfig, error) {
	// Load in a config file with the given name using the json parser.
	filename := "config.json"
	if err := k.Load(file.Provider(filename), json.Parser()); err != nil {
		return nil, err
	}

	// Load in the environment variables that correspond to what we want
	k.Load(env.Provider("" /* Prefix */, "." /* Delimeter */, func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil /* Parser */)

	config := &StructConfig{}
	err := k.Unmarshal("" /**/, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (s *StructConfig) GetString(name string) (string, error) {
	value := k.String(strings.ToLower(name))
	if value == "" {
		return value, fmt.Errorf("Could not find value %v", name)
	}
	return value, nil
}

func (s *StructConfig) GetInt(name string) (int, error) {
	value := k.Int(strings.ToLower(name))
	if value == 0 {
		return value, fmt.Errorf("Could not find value %v", name)
	}
	return value, nil
}

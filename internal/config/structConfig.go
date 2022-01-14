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

type StructConfig struct {
	filename string
	koanf    *koanf.Koanf
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

// NewConfig takes in a filename and unmarshals it into a config struct.
func ProvideConfig(k *koanf.Koanf, filesystem fs.FS) (*StructConfig, error) {
	config := &StructConfig{
		filename: "config.json",
		koanf:    k,
	}

	// Check if the file exists, if not, just use the environment variables.
	// Load in the config file using the json parser to read the variables into the config struct
	if _, err := os.Stat(config.filename); err == nil {
		if err := k.Load(fsys.Provider(filesystem, config.filename), json.Parser()); err != nil {
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

	err = k.Unmarshal("" /* path */, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func ProvideKoanf() *koanf.Koanf {
	return koanf.New(".")
}

func ProvideFilesystem() fs.FS {
	return os.DirFS(".")
}

func (s *StructConfig) GetString(name string) (string, error) {
	value := s.koanf.String(strings.ToLower(name))
	if value == "" {
		return "", fmt.Errorf("Could not find value %v", name)
	}
	return value, nil
}

func (s *StructConfig) GetInt(name string) (int, error) {
	value := s.koanf.Int(strings.ToLower(name))
	if value == 0 {
		return 0, fmt.Errorf("Could not find value %v", name)
	}
	return value, nil
}

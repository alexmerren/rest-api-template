package config

import (
	"io/fs"
	"os"
	"rest-api-template/internal/domain/entities"
	"strconv"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	fsys "github.com/knadh/koanf/providers/fs"
)

var k = koanf.New(".")

type Configuration struct {
	m map[string]interface{}
}

func NewConfiguration(filename string, filesystem fs.FS) *Configuration {
	if _, err := filesystem.Open(filename); err == nil {
		if err := k.Load(fsys.Provider(filesystem, filename), yaml.Parser()); err != nil {
			return nil
		}
	}

	// Load in the environment variables that correspond to what we want in the config struct
	err := k.Load(env.Provider("REST_" /* Prefix */, "." /* Delimeter */, func(s string) string {
		return strings.Replace(strings.ToLower(strings.TrimPrefix(s, "REST_")), "_", ".", -1)
	}), nil /* Parser */)
	if err != nil {
		return nil
	}

	return &Configuration{
		m: k.All(),
	}
}

func (c *Configuration) GetString(name string) (string, error) {
	value, ok := c.m[strings.ToLower(name)].(string)
	if !ok {
		return "", entities.NewNotFoundError("there was an error", nil)
	}
	return value, nil
}

func (c *Configuration) GetInt(name string) (int, error) {
	rawValue := c.m[strings.ToLower(name)]
	switch rawValue.(type) {
	case int:
		return rawValue.(int), nil
	case string:
		return strconv.Atoi(rawValue.(string))
	default:
		return 0, entities.NewNotFoundError("there was an error", nil)
	}
}

func NewFilesystem() fs.FS {
	return os.DirFS(".")
}

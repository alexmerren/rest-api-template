package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	values map[string]interface{}
}

// ReadInConfig takes in a filename and unmarshals it into a config struct.
func ReadInConfig() (*Configuration, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config := &Configuration{
		values: make(map[string]interface{}),
	}
	if err = json.Unmarshal([]byte(raw), &config.values); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Configuration) GetString(name string) (string, error) {
	value, ok := c.values[name].(string)
	if !ok {
		return "", fmt.Errorf("could not find config value: %v", name)
	}
	return value, nil
}

func (c *Configuration) GetInt(name string) (int, error) {
	value, ok := c.values[name].(float64)
	if !ok {
		return 0, fmt.Errorf("could not find config value: %v", name)
	}
	return int(value), nil
}

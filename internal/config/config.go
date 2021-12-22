package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Port   int    `json:"Port"`
	Host   string `json:"Host"`
	Logger ConfigurationLogger
}

type ConfigurationLogger struct {
	Level    string `json:"Level"`
	Encoding string `json:"Encoding"`
}

func ReadInConfig(filename string) (*Configuration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Configuration{}
	jsonDecoder := json.NewDecoder(file)
	if err = jsonDecoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

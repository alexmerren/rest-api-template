package config

import (
	"encoding/json"
	"os"
)

var config *Configuration

type Configuration struct {
	Port     int    `json:"Port"`
	Host     string `json:"Host"`
	Logger   ConfigurationLogger
	Database ConfigurationDatabase
}

type ConfigurationLogger struct {
	Level    string `json:"Level"`
	Encoding string `json:"Encoding"`
}

type ConfigurationDatabase struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
	Port     int    `json:"Port"`
}

func ReadInConfig(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonDecoder := json.NewDecoder(file)
	if err = jsonDecoder.Decode(&config); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Configuration {
	return config
}

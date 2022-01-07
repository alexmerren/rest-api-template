package config

type ConfigInterface interface {
	GetString(string) (string, error)
	GetInt(string) (int, error)
}

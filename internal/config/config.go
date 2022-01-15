package config

type Config interface {
	GetString(string) (string, error)
	GetInt(string) (int, error)
}

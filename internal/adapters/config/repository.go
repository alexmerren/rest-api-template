package config

type Configurator interface {
	GetString(name string) (string, error)
	GetInt(name string) (int, error)
}

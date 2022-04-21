package config

type Retriever interface {
	GetString(name string) (string, error)
	GetInt(name string) (int, error)
}

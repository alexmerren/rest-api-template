package config

type Configuration struct {
	m map[string]string
}

func NewConfiguration() Retriever {
	return &Configuration{}
}

func (c *Configuration) GetString(name string) (string, error) {
	return "", nil
}

func (c *Configuration) GetInt(name string) (int, error) {
	return 0, nil
}

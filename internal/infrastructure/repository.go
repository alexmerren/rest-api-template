package infrastructure

type Server interface {
	Start() error
	Stop()
}

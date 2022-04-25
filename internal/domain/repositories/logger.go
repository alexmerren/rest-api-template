package repositories

type Logger interface {
	Info(msg interface{})
	Error(msg interface{})
	Debug(msg interface{})
	WithField(name string, msg interface{}) Logger
	Cleanup()
}

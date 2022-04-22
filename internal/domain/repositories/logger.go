package repositories

type Logger interface {
	Info(msg interface{})
	Error(msg interface{})
	Debug(msg interface{})
	Cleanup()
}

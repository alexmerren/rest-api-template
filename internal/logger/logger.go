package logger

type Logger interface {
	Info(msg interface{})
	Debug(msg interface{})
	Error(msg interface{})
}

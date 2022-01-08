package logger

type LoggerInterface interface {
	Info(msg interface{})
	Debug(msg interface{})
	Error(msg interface{})
}

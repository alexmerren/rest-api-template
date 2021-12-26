package logger

type LoggerInterface interface {
	Error(msg interface{})
	Info(msg interface{})
}

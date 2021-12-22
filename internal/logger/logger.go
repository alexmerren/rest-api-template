package logger

type LoggerInterface interface {
	Error(msg string)
	Info(msg string)
}

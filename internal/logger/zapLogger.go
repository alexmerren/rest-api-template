package logger

import (
	"golang-api-template/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() (*ZapLogger, error) {
	config := config.GetConfig()
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(config.Logger.Level)); err != nil {
		return nil, err
	}

	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Encoding:         config.Logger.Encoding,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     "\n",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		},
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{
		logger: logger.Sugar(),
	}, nil
}

func (z *ZapLogger) Error(msg interface{}) {
	z.logger.Error(msg)
}

func (z *ZapLogger) Info(msg interface{}) {
	z.logger.Info(msg)
}

package logger

import (
	"rest-api-template/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

// NewZapLogger returns a new properly configured zap logger
func ProvideLogger(config config.Config) (*ZapLogger, error) {
	levelString, levelErr := config.GetString("logger.level")
	if levelErr != nil {
		return nil, levelErr
	}

	encoding, encodingErr := config.GetString("logger.encoding")
	if encodingErr != nil {
		return nil, encodingErr
	}

	var level zapcore.Level
	if err := level.UnmarshalText([]byte(levelString)); err != nil {
		return nil, err
	}

	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Encoding:         encoding,
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

func (z *ZapLogger) Debug(msg interface{}) {
	z.logger.Debug(msg)
}

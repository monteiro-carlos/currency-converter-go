package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Zap *zap.Logger
}

func NewLogger() (*Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return &Logger{
		Zap: logger,
	}, nil
}

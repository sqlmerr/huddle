package logger

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger

	// file *os.File
}

var LoggerKey = "logger"

func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(LoggerKey).(*Logger)
	if !ok {
		panic("no logger in context")
	}
	return logger
}

func New(config Config) (*Logger, error) {
	lvl := zap.NewAtomicLevel()
	if err := lvl.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, fmt.Errorf("unmarshal log level: %w", err)
	}

	// TODO: log file

	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")

	zapEncoder := zapcore.NewConsoleEncoder(zapConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), lvl),
		// TODO: file write
	)

	zapLogger := zap.New(core, zap.AddCaller())
	return &Logger{zapLogger}, nil
}

func (l *Logger) Close() {
	// TODO: file close
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{Logger: l.Logger.With(fields...)}
}

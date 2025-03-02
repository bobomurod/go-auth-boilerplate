package baseLogger

import (
	"github.com/bobomurod/go-auth-bolilerplate/internal/core/ports"
	"log/slog"
	"os"
)

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger() ports.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))
	return &SlogLogger{logger: logger}
}

func (l *SlogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *SlogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *SlogLogger) With(args ...any) ports.Logger {
	return &SlogLogger{
		logger: l.logger.With(args...),
	}
}

package logger

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"

	"time"
)

type Logger interface {
	GetZapLogger() *zap.Logger

	Named(s string) *logger

	With(fields ...zap.Field) *logger

	Debug(ctx context.Context, msg string, fields ...zap.Field)

	Info(ctx context.Context, msg string, fields ...zap.Field)

	Warn(ctx context.Context, msg string, fields ...zap.Field)

	Error(ctx context.Context, msg string, fields ...zap.Field)

	Panic(ctx context.Context, msg string, fields ...zap.Field)

	Fatal(ctx context.Context, msg string, fields ...zap.Field)

	Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{})

	extract(ctx context.Context) []zap.Field
}

type logger struct {
	logger *zap.Logger
}

func New(l *zap.Logger) Logger {
	return &logger{l}
}

func (l *logger) GetZapLogger() *zap.Logger {
	return l.logger
}

func (l *logger) Named(s string) *logger {
	l2 := l.logger.Named(s)
	return &logger{l2}
}

func (l *logger) With(fields ...zap.Field) *logger {
	l2 := l.logger.With(fields...)
	return &logger{l2}
}

func (l *logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Debug(msg, fields...)
}

func (l *logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Info(msg, fields...)
}

func (l *logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Warn(msg, fields...)
}

func (l *logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Error(msg, fields...)
}

func (l *logger) Panic(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Panic(msg, fields...)
}

func (l *logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Fatal(msg, fields...)
}

func (l *logger) extract(ctx context.Context) []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.String("time", time.Now().Format(time.RFC3339)))

	if reqID, ok := ctx.Value("x-request-id").(string); ok {
		fields = append(fields, zap.String("x-request-id", reqID))
	}

	if hitTime, ok := ctx.Value("request-start-time").(time.Time); ok {
		fields = append(fields, zap.Float64("time-since-request", float64(time.Since(hitTime).Milliseconds())))
	}

	return fields
}

func (l *logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	var fields []zap.Field
	// since the logger already has a `time` field, we have to remap the pgx time
	data["pgx_time"] = data["time"]
	delete(data, "time")
	for k, v := range data {
		fields = append(fields, zap.Any(k, v))
	}
	switch level {
	case pgx.LogLevelInfo:
		l.Info(ctx, msg, fields...)
	case pgx.LogLevelWarn:
		l.Warn(ctx, msg, fields...)
	case pgx.LogLevelError:
		l.Error(ctx, msg, fields...)
	default:
		l.Debug(ctx, msg, fields...)
	}
}

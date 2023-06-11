package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(service string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": service,
	}

	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	return log.Sugar(), nil
}

type logCtxKeyType struct{}

var logCtxKey logCtxKeyType = logCtxKeyType{}

func WithLogContext(ctx context.Context, log *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, logCtxKey, log)
}

func LogFromContext(ctx context.Context) *zap.SugaredLogger {
	log, ok := ctx.Value(logCtxKey).(*zap.SugaredLogger)
	if ok && log != nil {
		return log
	}
	return nil
}

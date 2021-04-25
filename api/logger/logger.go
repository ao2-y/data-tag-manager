package logger

import (
	"context"
	"go.uber.org/zap"
)

const Key = "contextLogger"

var appLogger *zap.Logger

// InitApplicationLogger Rootになるロガー
func InitApplicationLogger() *zap.Logger {
	if appLogger != nil {
		return appLogger
	}
	var err error
	// FIXME ログの設定何とかする
	appLogger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return appLogger
}

func Ctx(ctx context.Context) *zap.Logger {
	if ctxLogger, ok := ctx.Value(Key).(*zap.Logger); ok {
		return ctxLogger
	}
	return appLogger
}

func With(ctx context.Context, fields ...zap.Field) context.Context {
	log := Ctx(ctx)
	log = log.With(fields...)
	return context.WithValue(ctx, Key, log)
}

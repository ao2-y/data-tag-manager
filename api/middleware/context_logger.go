package middleware

import (
	"ao2-y/data-tag-manager/logger"
	"context"
	"github.com/dgryski/trifles/uuid"
	"go.uber.org/zap"
	"net/http"
)

func ContextLogger(appLogger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctxLogger := appLogger.With(zap.String("RequestID", uuid.UUIDv4()))
			ctx = context.WithValue(ctx, logger.Key, ctxLogger)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

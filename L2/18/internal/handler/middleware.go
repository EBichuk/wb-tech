package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func MiddlewareLogging(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			timeStart := time.Now()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			status := ww.Status()

			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", status),
				slog.String("duration", time.Since(timeStart).String()),
			)

			msg := fmt.Sprintf("request %s", r.URL.Path)

			switch status {
			case 500:
				entry.Error(msg)
			case 400, 503:
				entry.Warn(msg)
			default:
				entry.Info(msg)
			}

		})
	}
}

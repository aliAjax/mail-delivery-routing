package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func RequestLog(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			if log != nil {
				log.Info("http request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start).String())
			}
		})
	}
}
func Recover(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if v := recover(); v != nil {
					if log != nil {
						log.Error("panic", "value", v)
					}
					http.Error(w, "internal error", 500)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
func Limit(max int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ContentLength > max {
				http.Error(w, "payload too large", 413)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

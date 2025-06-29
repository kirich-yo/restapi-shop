package middleware

import (
	"net/http"
	"log/slog"
	"time"
	"fmt"
)

func Logger(next http.Handler, args ...interface{}) http.Handler {
	logger := args[0].(*slog.Logger)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("New request:",
			slog.String("remoteAddr", r.RemoteAddr),
			slog.String("method", r.Method),
			slog.String("path", r.RequestURI),
			slog.String("userAgent", r.Header.Get("User-Agent")),
			slog.String("contentType", r.Header.Get("Content-Type")),
		)

		ww := NewResponseWriterWrapper(w)
		t := time.Now()

		next.ServeHTTP(ww, r)

		logger.Info("Request completed:",
			slog.String("status", fmt.Sprintf("%d %s", ww.StatusCode, http.StatusText(ww.StatusCode))),
			slog.String("timeElapsed", time.Since(t).String()),
		)
	})
}

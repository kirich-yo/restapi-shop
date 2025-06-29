package middleware

import (
	"net/http"
	"log/slog"
	"runtime/debug"
	"fmt"
)

func Recoverer(next http.Handler, args ...interface{}) http.Handler {
	logger := args[0].(*slog.Logger)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				recText := fmt.Sprintf("%v", rec)

				logger.Error(recText)
				logger.Debug(fmt.Sprintf("Stacktrace: %s", debug.Stack()))
				http.Error(w, recText, http.StatusInternalServerError)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

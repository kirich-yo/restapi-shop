package middleware

import (
	"net/http"
	"fmt"
)

func Recoverer(next http.Handler, args ...interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				http.Error(w, fmt.Sprintf("%v", rec), http.StatusInternalServerError)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

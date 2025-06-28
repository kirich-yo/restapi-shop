package middleware

import (
	"net/http"
	"strings"
	"context"

	"restapi-sportshop/configs"
	"restapi-sportshop/pkg/jwt"
)

const (
	ContextUsernameKey = "username"
)

func Authorization(next http.Handler, args ...interface{}) http.Handler {
	cfg := args[0].(*configs.Config)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Invalid auth token", http.StatusUnauthorized)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		j := jwt.NewJWT(cfg.AuthConfig.Secret)

		data, err := j.Parse(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUsernameKey, data.UserID)
		ctxReq := r.WithContext(ctx)

		next.ServeHTTP(w, ctxReq)
	})
}

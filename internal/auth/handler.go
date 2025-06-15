package auth

import (
	"net/http"
	"log/slog"

	"restapi-sportshop/pkg/req"
	"restapi-sportshop/pkg/middleware"
)

type AuthHandler struct {
	*AuthService
	*slog.Logger
}

type AuthHandlerDeps struct {
	*AuthService
	*slog.Logger
}

func NewAuthHandler(smux *http.ServeMux, deps AuthHandlerDeps) *AuthHandler {
	handler := &AuthHandler{
		AuthService: deps.AuthService,
		Logger: deps.Logger,
	} 

	chain := middleware.Chain(
                middleware.MiddlewareWithArgs{
                        First: middleware.Logger,
                        Second: []interface{}{handler.Logger},
                },
                middleware.MiddlewareWithArgs{
                        First: middleware.CORS,
                },
        )

	smux.Handle("POST /auth/login", chain(handler.Login()))
	smux.Handle("POST /auth/register", chain(handler.Register()))
	smux.Handle("POST /auth/refresh", chain(handler.Refresh()))

	return handler
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := req.HandleBody[RegisterRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.AuthService.Register(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *AuthHandler) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

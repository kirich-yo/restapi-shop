package auth

import (
	"net/http"
	"log/slog"
	"io"

	"restapi-sportshop/configs"
	"restapi-sportshop/pkg/jwt"
	"restapi-sportshop/pkg/req"
	"restapi-sportshop/pkg/middleware"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
	*slog.Logger
}

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
	*slog.Logger
}

func NewAuthHandler(smux *http.ServeMux, deps AuthHandlerDeps) *AuthHandler {
	handler := &AuthHandler{
		Config: deps.Config,
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
		defer r.Body.Close()

		body, err := req.HandleBody[LoginRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, err := handler.AuthService.Login(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token := jwt.NewJWT(handler.Config.AuthConfig.Secret)

		s, err := token.Create(userID, handler.Config.AuthConfig.TokenLifetime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "plain/text; encoding=utf-8")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, s)
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

		userID, err := handler.AuthService.Register(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token := jwt.NewJWT(handler.Config.AuthConfig.Secret)

		s, err := token.Create(userID, handler.Config.AuthConfig.TokenLifetime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "plain/text; encoding=utf-8")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, s)
	}
}

func (handler *AuthHandler) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

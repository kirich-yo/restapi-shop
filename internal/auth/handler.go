package auth

import (
	"net/http"
)

type AuthHandler struct {
}

type AuthHandlerDeps struct {
}

func NewAuthHandler(smux *http.ServeMux, deps AuthHandlerDeps) *AuthHandler {
	handler := &AuthHandler{} 

	smux.HandleFunc("POST /auth/login", handler.Login())
	smux.HandleFunc("POST /auth/register", handler.Register())
	smux.HandleFunc("POST /auth/refresh", handler.Refresh())

	return handler
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (handler *AuthHandler) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

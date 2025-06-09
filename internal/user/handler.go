package user

import (
	"net/http"
	"io"
)

type UserHandler struct {
}

type UserHandlerDeps struct {
}

func NewUserHandler(smux *http.ServeMux, deps UserHandlerDeps) *UserHandler {
	handler := &UserHandler{}

	smux.Handle("GET /user", handler.Create())

	return handler
}

func (u *UserHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; encoding=utf-8")

		w.WriteHeader(http.StatusCreated)

		io.WriteString(w, `{
  "id": 456,
  "username": "johndoe",
  "firstName": "John",
  "lastName": "Doe",
  "dateOfBirth": "2000-01-01",
  "photoURL": "https://cdn.sportshop.com/59b1f1ce-7299-4e9b-93c4-cb5b94641864.jpg"
}`)

	}
}

package user

import (
	"net/http"

	"restapi-sportshop/pkg/res"
	"restapi-sportshop/pkg/req"

	"github.com/davecgh/go-spew/spew"
)

type UserHandler struct {
}

type UserHandlerDeps struct {
}

func NewUserHandler(smux *http.ServeMux, deps UserHandlerDeps) *UserHandler {
	handler := &UserHandler{}

	smux.HandleFunc("GET /user/{username}", handler.Get())
	smux.HandleFunc("POST /user", handler.Create())
	smux.HandleFunc("DELETE /user/{username}", handler.Delete())
	smux.HandleFunc("GET /user/{username}/role", handler.GetRole())

	return handler
}

func (u *UserHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := UserResponse{
			ID: 456,
			Username: r.PathValue("username"),
			FirstName: "John",
			LastName: "Doe",
			DateOfBirth: "2000-01-01",
			PhotoURL: "https://cdn.sportshop.com/59b1f1ce-7299-4e9b-93c4-cb5b94641864.jpg",
		}

		res.WriteDefault(w, http.StatusCreated, &data, r.Header)
	}
}

func (u *UserHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := req.HandleBody[UserRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		spew.Dump(body)

		w.WriteHeader(http.StatusCreated)
	}
}

func (u *UserHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (u *UserHandler) GetRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

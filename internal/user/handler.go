package user

import (
	"net/http"
	"fmt"
	"io"

	"restapi-sportshop/pkg/res"
)

type UserHandler struct {
}

type UserHandlerDeps struct {
}

func NewUserHandler(smux *http.ServeMux, deps UserHandlerDeps) *UserHandler {
	handler := &UserHandler{}

	smux.Handle("GET /user/{username}", handler.Create())

	return handler
}

func (u *UserHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Cannot read body", http.StatusInternalServerError)
			return
		}
		fmt.Println(string(body))

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

package user

import (
	"errors"
	"time"
	"net/http"

	"restapi-sportshop/pkg/res"
	"restapi-sportshop/pkg/req"

	"gorm.io/gorm"
	"github.com/davecgh/go-spew/spew"
)

type UserHandler struct {
	*UserRepository
}

type UserHandlerDeps struct {
	*UserRepository
}

func NewUserHandler(smux *http.ServeMux, deps UserHandlerDeps) *UserHandler {
	handler := &UserHandler{
		UserRepository: deps.UserRepository,
	}

	smux.HandleFunc("GET /user/{username}", handler.Get())
	smux.HandleFunc("POST /user", handler.Create())
	smux.HandleFunc("DELETE /user/{username}", handler.Delete())
	smux.HandleFunc("GET /user/{username}/role", handler.GetRole())

	return handler
}

func (handler *UserHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.PathValue("username")

                data, err := handler.UserRepository.GetByUsername(username)
                if errors.Is(err, gorm.ErrRecordNotFound) {
                        http.Error(w, err.Error(), http.StatusNotFound)
                        return
                }
                if err != nil {
                        http.Error(w, err.Error(), http.StatusBadRequest)
                        return
                }

                body := &UserResponse{
                        ID: data.ID,
                        Username: data.Username,
			FirstName: data.FirstName,
			LastName: data.LastName,
			DateOfBirth: time.Time(data.DateOfBirth).Format(time.DateOnly),
                        PhotoURL: data.PhotoURL,
                }

                res.WriteDefault(w, http.StatusOK, body, r.Header)
	}
}

func (handler *UserHandler) Create() http.HandlerFunc {
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

func (handler *UserHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (handler *UserHandler) GetRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

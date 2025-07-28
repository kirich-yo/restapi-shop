package user

import (
	"errors"
	"time"
	"net/http"
	"strconv"

	"restapi-shop/pkg/res"
	"restapi-shop/pkg/req"

	"gorm.io/gorm"
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
	smux.HandleFunc("PATCH /user/{username}", handler.Update())
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

func (handler *UserHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := req.HandleBody[UserUpdateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_ = body

		w.WriteHeader(http.StatusOK)
	}
}

func (handler *UserHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseUint(r.PathValue("userID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.UserRepository.Delete(uint(userID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (handler *UserHandler) GetRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

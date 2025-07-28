package review

import (
	"net/http"
	"strconv"
	"log/slog"

	"restapi-shop/configs"
	"restapi-shop/pkg/req"
	"restapi-shop/pkg/res"
	"restapi-shop/pkg/middleware"
)

type ReviewHandlerDeps struct {
	*ReviewService
	*configs.Config
	*slog.Logger
}

type ReviewHandler struct {
	*ReviewService
	*configs.Config
	*slog.Logger
}

func NewReviewHandler(smux *http.ServeMux, deps ReviewHandlerDeps) *ReviewHandler {
	handler := &ReviewHandler{
		ReviewService: deps.ReviewService,
		Config: deps.Config,
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
		middleware.MiddlewareWithArgs{
			First: middleware.Authorization,
			Second: []interface{}{handler.Config},
		},
		middleware.MiddlewareWithArgs{
			First: middleware.Recoverer,
			Second: []interface{}{handler.Logger},
		},
        )

	smux.Handle("GET /review/{reviewID}", chain(handler.Get()))
	smux.Handle("POST /review", chain(handler.Create()))
	smux.Handle("PATCH /review/{reviewID}", chain(handler.Update()))
	smux.Handle("DELETE /review/{reviewID}", chain(handler.Delete()))

	return handler
}

func (handler *ReviewHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		authUserID, ok := r.Context().Value(middleware.ContextUsernameKey).(uint)
		if !ok {
			http.Error(w, "Auth username is not uint", http.StatusUnauthorized)
			return
		}

		body, err := req.HandleBody[ReviewCreateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		review := NewReview((*ReviewRequest)(body), authUserID)

		review, err = handler.ReviewService.Create(review)
		switch err {
		case nil:
			 break
		case ErrFKViolated:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resBody := NewReviewResponse(review)

		res.WriteDefault(w, http.StatusCreated, resBody, r.Header)
	}
}

func (handler *ReviewHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		review_id, err := strconv.ParseUint(r.PathValue("reviewID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		review, err := handler.ReviewService.Get(uint(review_id))
		switch err {
		case nil:
			 break
		case ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body := NewReviewResponse(review)

		res.WriteDefault(w, http.StatusOK, body, r.Header)
	}
}

func (handler *ReviewHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		authUserID, ok := r.Context().Value(middleware.ContextUsernameKey).(uint)
		if !ok {
			http.Error(w, "Auth username is not uint", http.StatusUnauthorized)
			return
		}

		reviewID, err := strconv.ParseUint(r.PathValue("reviewID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, err := req.HandleBody[ReviewUpdateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		review := NewReview((*ReviewRequest)(body), authUserID)
		review.ID = uint(reviewID)

		updatedReview, err := handler.ReviewService.Update(review, authUserID)
		switch err {
		case nil:
			 break
		case ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case ErrNoPermission:
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resBody := NewReviewResponse(updatedReview)

		res.WriteDefault(w, http.StatusOK, resBody, r.Header)
	}
}

func (handler *ReviewHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authUserID, ok := r.Context().Value(middleware.ContextUsernameKey).(uint)
		if !ok {
			http.Error(w, "Auth username is not string", http.StatusUnauthorized)
			return
		}

		reviewID, err := strconv.ParseUint(r.PathValue("reviewID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.ReviewService.Delete(uint(reviewID), authUserID)
		switch err {
		case nil:
			 break
		case ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case ErrNoPermission:
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

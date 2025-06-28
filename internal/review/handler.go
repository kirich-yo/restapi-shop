package review

import (
	"net/http"
	"strconv"
	"log/slog"

	"restapi-sportshop/configs"
	"restapi-sportshop/pkg/req"
	"restapi-sportshop/pkg/res"
	"restapi-sportshop/pkg/middleware"
)

type ReviewHandlerDeps struct {
	*ReviewRepository
	*configs.Config
	*slog.Logger
}

type ReviewHandler struct {
	*ReviewRepository
	*configs.Config
	*slog.Logger
}

func NewReviewHandler(smux *http.ServeMux, deps ReviewHandlerDeps) *ReviewHandler {
	handler := &ReviewHandler{
		ReviewRepository: deps.ReviewRepository,
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
			http.Error(w, "Auth username is not string", http.StatusUnauthorized)
			return
		}

		body, err := req.HandleBody[ReviewCreateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		review := NewReview((*ReviewRequest)(body), authUserID)

		_, err = handler.ReviewRepository.Create(review)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *ReviewHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		review_id, err := strconv.ParseUint(r.PathValue("reviewID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		review, err := handler.ReviewRepository.Get(uint(review_id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body := &ReviewResponse{
			ID: review.ID,
			UserID: review.UserID,
			ItemID: review.ItemID,
			Rating: review.Rating,
			Advantages: review.Advantages,
			Disadvantages: review.Disadvantages,
			Description: review.Description,
		}

		res.WriteDefault(w, http.StatusOK, body, r.Header)
	}
}

func (handler *ReviewHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		review, err := handler.ReviewRepository.Get(uint(reviewID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if review.UserID != authUserID {
			http.Error(w, "You have not enough permissions to delete the review.", http.StatusUnauthorized)
			return
		}

		err = handler.ReviewRepository.Delete(uint(reviewID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

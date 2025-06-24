package review

import (
	"net/http"
	"strconv"

	_ "restapi-sportshop/pkg/req"
	_ "restapi-sportshop/pkg/res"
)

type ReviewHandlerDeps struct {
	*ReviewRepository
}

type ReviewHandler struct {
	*ReviewRepository
}

func NewReviewHandler(smux *http.ServeMux, deps ReviewHandlerDeps) *ReviewHandler {
	handler := &ReviewHandler{
		ReviewRepository: deps.ReviewRepository,
	}

	smux.HandleFunc("GET /review/{reviewID}", handler.Get())

	return handler
}

func (handler *ReviewHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		review_id, err := strconv.ParseUint(r.PathValue("reviewID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		review, err := handler.ReviewRepository.Get(uint(review_id))
		_ = review
	}
}

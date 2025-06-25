package review

import (
	"net/http"
	"strconv"
	"fmt"

	"restapi-sportshop/configs"
	_ "restapi-sportshop/pkg/req"
	_ "restapi-sportshop/pkg/res"
	"restapi-sportshop/pkg/middleware"
)

type ReviewHandlerDeps struct {
	*ReviewRepository
	*configs.Config
}

type ReviewHandler struct {
	*ReviewRepository
	*configs.Config
}

func NewReviewHandler(smux *http.ServeMux, deps ReviewHandlerDeps) *ReviewHandler {
	handler := &ReviewHandler{
		ReviewRepository: deps.ReviewRepository,
		Config: deps.Config,
	}

	smux.Handle("GET /review/{reviewID}", middleware.Authorization(handler.Get(), handler.Config.AuthConfig.Secret))

	return handler
}

func (handler *ReviewHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authUsername, ok := r.Context().Value(middleware.ContextUsernameKey).(string)
		if !ok {
			http.Error(w, "Auth username is not string", http.StatusUnauthorized)
			return
		}
		fmt.Println(authUsername)

		review_id, err := strconv.ParseUint(r.PathValue("reviewID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		review, err := handler.ReviewRepository.Get(uint(review_id))
		_ = review
	}
}

package item

import (
	"net/http"
	"strconv"
	"errors"
	"log/slog"

	"restapi-shop/pkg/res"
	"restapi-shop/pkg/req"
	"restapi-shop/pkg/middleware"

	"gorm.io/gorm"
)

type ItemHandler struct {
	*ItemRepository
	*slog.Logger
}

type ItemHandlerDeps struct {
	*ItemRepository
	*slog.Logger
}

func NewItemHandler(smux *http.ServeMux, deps ItemHandlerDeps) *ItemHandler {
	handler := &ItemHandler{
		ItemRepository: deps.ItemRepository,
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
	)

	smux.Handle("GET /item/{itemID}", chain(handler.Get()))
	smux.Handle("POST /item", chain(handler.Create()))
	smux.Handle("PATCH /item/{itemID}", chain(handler.Update()))
	smux.Handle("DELETE /item/{itemID}", chain(handler.Delete()))

	return handler
}

func (handler *ItemHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		item_id, err := strconv.ParseUint(r.PathValue("itemID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := handler.ItemRepository.Get(uint(item_id))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body := &ItemResponse{
			ID: data.ID,
			Name: data.Name,
			Price: data.Price,
			SalePrice: data.SalePrice,
			PhotoURL: data.PhotoURL,
		}

		res.WriteDefault(w, http.StatusOK, body, r.Header)
	}
}

func (handler *ItemHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := req.HandleBody[ItemCreateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := NewItem((*ItemRequest)(body))

		_, err = handler.ItemRepository.Create(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *ItemHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		itemID, err := strconv.ParseUint(r.PathValue("itemID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.ItemRepository.Delete(uint(itemID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (handler *ItemHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		item_id, err := strconv.ParseUint(r.PathValue("itemID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		found, err := handler.ItemRepository.IsExist(uint(item_id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !found {
			http.Error(w, "record not found", http.StatusNotFound)
			return
		}

		body, err := req.HandleBody[ItemUpdateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := NewItem((*ItemRequest)(body))
		data.ID = uint(item_id)

		item, err := handler.ItemRepository.Update(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.WriteDefault(w, http.StatusOK, item, r.Header)
	}
}

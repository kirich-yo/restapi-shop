package item

import (
	"net/http"
	"strconv"
	"fmt"

	"restapi-sportshop/pkg/res"
	"restapi-sportshop/pkg/req"

	"github.com/davecgh/go-spew/spew"
)

type ItemHandler struct {
	*ItemRepository
}

type ItemHandlerDeps struct {
	*ItemRepository
}

func NewItemHandler(smux *http.ServeMux, deps ItemHandlerDeps) *ItemHandler {
	handler := &ItemHandler{
		ItemRepository: deps.ItemRepository,
	}

	smux.Handle("GET /item/{itemID}", handler.Get())
	smux.Handle("POST /item", handler.Create())
	smux.Handle("PATCH /item/{itemID}", handler.Update())
	smux.Handle("DELETE /item/{itemID}", handler.Delete())

	return handler
}

func (handler *ItemHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		item_id, err := strconv.ParseUint(r.PathValue("itemID"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var count int64 = 0
		_ = handler.ItemRepository.Count(&count)
		fmt.Printf("Count: %d\n", count)

		data, err := handler.ItemRepository.Get(uint(item_id))
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

		res.WriteDefault(w, http.StatusCreated, body, r.Header)
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
		item_id, err := strconv.Atoi(r.PathValue("itemID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if item_id < 1 {
			http.Error(w, "ID number cannot be negative or zero.", http.StatusBadRequest)
			return
		}

		err = handler.ItemRepository.Delete(item_id)
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

package item

import (
	"net/http"
	"strconv"

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

	return handler
}

func (handler *ItemHandler) Get() http.HandlerFunc {
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

		data, err := handler.ItemRepository.Get(item_id)
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

		body, err := req.HandleBody[ItemRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := NewItem(body)
		spew.Dump(data)

		_, err = handler.ItemRepository.Create(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

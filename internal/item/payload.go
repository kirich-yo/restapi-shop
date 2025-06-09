package item

import (
	"encoding/xml"
)

type ItemRequest struct {
	XMLName xml.Name `json:"-" xml:"item"`
	Name string `json:"name" xml:"name" validate:"required"`
	Price float64 `json:"price" xml:"price" validate:"required"`
	SalePrice float64 `json:"salePrice" xml:"salePrice"`
	PhotoURL string `json:"photoURL" xml:"photoURL" validate:"http_url"`
}

type ItemResponse struct {
	XMLName xml.Name `json:"-" xml:"item"`
	ID uint `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Price float64 `json:"price" xml:"price"`
	SalePrice float64 `json:"salePrice" xml:"salePrice"`
	PhotoURL string `json:"photoURL" xml:"photoURL"`
}

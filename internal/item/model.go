package item

import (
	"gorm.io/gorm"

	"restapi-sportshop/internal/review"
)

type Item struct {
	gorm.Model
	Name string `json:"name"`
	Price float64 `json:"price"`
	SalePrice float64 `json:"sale_price"`
	PhotoURL string `json:"photo_url"`
	Reviews []review.Review `gorm:"constraints:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewItem(r *ItemRequest) *Item {
	return &Item{
		Name: r.Name,
		Price: r.Price,
		SalePrice: r.SalePrice,
		PhotoURL: r.PhotoURL,
	}
}

package cart

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint `json:"user_id"`
	ItemID uint `json:"item_id"`
	Quantity uint `json:"quantity"`
}

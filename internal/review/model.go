package review

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID uint `json:"user_id"`
	ItemID uint `json:"item_id"`
	Rating uint `json:"rating"`
	Advantages string `json:"advantages"`
	Disadvantages string `json:"disadvantages"`
	Description string `json:"description"`
}

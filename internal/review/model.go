package review

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID          uint     `json:"user_id"`
	ItemID          uint     `json:"item_id"`
	Rating          uint     `json:"rating"`
	Advantages      string   `json:"advantages"`
	Disadvantages   string   `json:"disadvantages"`
	Description     string   `json:"description"`
}

func NewReview(r *ReviewRequest, userID uint) *Review {
	return &Review{
		UserID:        userID,
		ItemID:        r.ItemID,
		Rating:        r.Rating,
		Advantages:    r.Advantages,
		Disadvantages: r.Disadvantages,
		Description:   r.Description,
	}
}

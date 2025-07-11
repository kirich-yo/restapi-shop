package order

import (
	"gorm.io/gorm"
	"gorm.io/gorm/datatypes"
)

type OrderStatus string

const (
	OrderPlaced OrderStatus = "placed"
	OrderApproved = "approved"
	OrderDelayed = "delayed"
	OrderDelivered = "delivered"
)

type Order struct {
	gorm.Model
	UserID uint `json:"user_id"`
	ItemID uint `json:"item_id"`
	Amount uint `json:"amount"`
	ShipDate datatypes.Date `json:"ship_date"`
	Status OrderStatus `json:"status"`
}

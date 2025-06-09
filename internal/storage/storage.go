package storage

import (
	"time"
	"errors"
)

type OrderStatus int

const (
        OrderPlaced OrderStatus = iota
        OrderApproved
        OrderDelayed
        OrderDelivered
)

func (os *OrderStatus) String() string {
	switch *os {
	case OrderPlaced:
		return "placed"
	case OrderApproved:
		return "approved"
	case OrderDelayed:
		return "delayed"
	case OrderDelivered:
		return "delivered"
	default:
		return "-"
	}
}

var (
	ErrItemExists = errors.New("item exists")
	ErrItemNotFound = errors.New("item not found")
)

type Item struct {
        ID          int     `json:"id"`
        Name        string  `json:"name"`
        Price       float64 `json:"price"`
        SalePrice   float64 `json:"salePrice"`
        PhotoURL    string  `json:"photoURL"`
}

type User struct {
        ID            int
        Username      string
        FirstName     string
        LastName      string
        DateOfBirth   time.Time
        PhotoURL      string
	RoleID        int
        Password      string
}

type Review struct {
        ID             int     `json:"id"`
        UserID         int     `json:"userID"`
        ItemID         int     `json:"itemID"`
        Rating         int     `json:"rating"`
        Advantages     string  `json:"advantages"`
        Disadvantages  string  `json:"disadvantages"`
        Description    string  `json:"description"`
}

type Order struct {
        ID          int
        UserID      int
        ItemID      int
        Quantity    int
        ShipDate    time.Time
        Status      OrderStatus
        Complete    bool
}

type Role struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
}

type Warehouse struct {
        ID      int    `json:"id"`
        Address string `json:"address"`
}

type Category struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
}

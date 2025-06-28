package review

import (
	"encoding/xml"
)

type ReviewRequest struct {
	XMLName xml.Name
	ItemID uint
	Rating uint
	Advantages string
	Disadvantages string
	Description string
}

type ReviewCreateRequest struct {
	XMLName xml.Name `json:"-" xml:"review"`
	ItemID uint `json:"itemID" xml:"itemID" validate:"required"`
	Rating uint `json:"rating" xml:"rating" validate:"min=1,max=5"`
	Advantages string `json:"advantages" xml:"advantages"`
	Disadvantages string `json:"disadvantages" xml:"disadvantages"`
	Description string `json:"description" xml:"description"`
}

type ReviewUpdateRequest struct {
	XMLName xml.Name `json:"-" xml:"review"`
	ItemID uint `json:"itemID" xml:"itemID"`
	Rating uint `json:"rating" xml:"rating" validate:"omitempty,min=1,max=5"`
	Advantages string `json:"advantages" xml:"advantages"`
	Disadvantages string `json:"disadvantages" xml:"disadvantages"`
	Description string `json:"description" xml:"description"`
}

type ReviewResponse struct {
	XMLName xml.Name `json:"-" xml:"review"`
	ID uint `json:"id" xml:"id"`
	UserID uint `json:"userID" xml:"userID"`
	ItemID uint `json:"itemID" xml:"itemID"`
	Rating uint `json:"rating" xml:"rating"`
	Advantages string `json:"advantages" xml:"advantages"`
	Disadvantages string `json:"disadvantages" xml:"disadvantages"`
	Description string `json:"description" xml:"description"`
}

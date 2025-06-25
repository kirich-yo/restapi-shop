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
	ItemID uint `json:"itemID" xml:"itemID"`
	Rating uint `json:"rating" xml:"rating"`
	Advantages string `json:"advantages" xml:"advantages"`
	Disadvantages string `json:"disadvantages" xml:"disadvantages"`
	Description string `json:"description" xml:"description"`
}

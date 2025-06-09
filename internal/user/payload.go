package user

import (
	"encoding/xml"
)

type UserResponse struct {
	XMLName xml.Name `json:"-" xml:"user"`
	ID uint `json:"id" xml:"id"`
	Username string `json:"username" xml:"username"`
	FirstName string `json:"firstName" xml:"firstName"`
	LastName string `json:"lastName" xml:"lastName"`
	DateOfBirth string `json:"dateOfBirth" xml:"dateOfBirth"`
	PhotoURL string `json:"photoURL" xml:"photoURL"`
}

package user

import (
	"time"
	"encoding/xml"
)

type UserRequest struct {
	XMLName xml.Name
	Username string
	FirstName string
	LastName string
	DateOfBirth string
	PhotoURL string
}

type UserCreateRequest struct {
	XMLName xml.Name `json:"-" xml:"user"`
	Username string `json:"username" xml:"username" validate:"required,alphanum"`
	FirstName string `json:"firstName" xml:"firstName" validate:"required,alpha"`
	LastName string `json:"lastName" xml:"lastName" validate:"required,alpha"`
	DateOfBirth string `json:"dateOfBirth" xml:"dateOfBirth" validate:"required,datetime=2006-01-02"`
	PhotoURL string `json:"photoURL" xml:"photoURL" validate:"http_url"`
}

type UserUpdateRequest struct {
	XMLName xml.Name `json:"-" xml:"user"`
	Username string `json:"username" xml:"username" validate:"omitempty,alphanum"`
	FirstName string `json:"firstName" xml:"firstName" validate:"omitempty,alpha"`
	LastName string `json:"lastName" xml:"lastName" validate:"omitempty,alpha"`
	DateOfBirth string `json:"dateOfBirth" xml:"dateOfBirth" validate:"omitempty,datetime=2006-01-02"`
	PhotoURL string `json:"photoURL" xml:"photoURL" validate:"omitempty,http_url"`
}

type UserResponse struct {
	XMLName xml.Name `json:"-" xml:"user"`
	ID uint `json:"id" xml:"id"`
	Username string `json:"username" xml:"username"`
	FirstName string `json:"firstName" xml:"firstName"`
	LastName string `json:"lastName" xml:"lastName"`
	DateOfBirth string `json:"dateOfBirth" xml:"dateOfBirth"`
	PhotoURL string `json:"photoURL" xml:"photoURL"`
}

func NewUserResponse(data *User) *UserResponse {
	return &UserResponse{
		ID: data.ID,
		Username: data.Username,
		FirstName: data.FirstName,
		LastName: data.LastName,
		DateOfBirth: time.Time(data.DateOfBirth).Format(time.DateOnly),
		PhotoURL: data.PhotoURL,
	}
}

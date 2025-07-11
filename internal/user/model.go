package user

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/datatypes"

	"restapi-sportshop/internal/review"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DateOfBirth datatypes.Date `json:"date_of_birth"`
	PhotoURL string `json:"photo_url"`
	RoleID uint `json:"role_id"`
	Password string `json:"password"`
	Reviews []review.Review `gorm:"constraints:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewUser(data *UserRequest) (*User, error) {
	parsedDate, err := time.Parse(time.DateOnly, data.DateOfBirth)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: data.Username,
		FirstName: data.FirstName,
		LastName: data.LastName,
		DateOfBirth: datatypes.Date(parsedDate),
		PhotoURL: data.PhotoURL,
	}, nil
}

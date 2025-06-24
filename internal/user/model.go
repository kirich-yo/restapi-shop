package user

import (
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

/*
func NewUser(data *auth.RegisterRequest) *User {
	return nil
}
*/

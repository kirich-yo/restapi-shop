package auth

import (
	"golang.org/x/crypto/bcrypt"
	"time"

	"restapi-shop/internal/user"

	"gorm.io/datatypes"
)

type AuthService struct {
	*user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (srv *AuthService) Login(data *LoginRequest) (uint, error) {
	loginUser, err := srv.UserRepository.GetByUsername(data.Username)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(data.Password))
	if err != nil {
		return 0, err
	}

	return loginUser.ID, nil
}

func (srv *AuthService) Register(data *RegisterRequest) (uint, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	parsedDate, err := time.Parse(time.DateOnly, data.DateOfBirth)
	if err != nil {
		return 0, err
	}

	newUser := &user.User{
		Username: data.Username,
		FirstName: data.FirstName,
		LastName: data.LastName,
		DateOfBirth: datatypes.Date(parsedDate),
		PhotoURL: data.PhotoURL,
		RoleID: 1,
		Password: string(hashedPassword),
	}

	newUser, err = srv.UserRepository.Create(newUser)
	if err != nil {
		return 0, err
	}

	return newUser.ID, nil
}

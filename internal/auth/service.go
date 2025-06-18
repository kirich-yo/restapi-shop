package auth

import (
	"golang.org/x/crypto/bcrypt"
	"time"

	"restapi-sportshop/internal/user"

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

func (srv *AuthService) Login(data *LoginRequest) error {
	loginUser, err := srv.UserRepository.GetByUsername(data.Username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(data.Password))
	if err != nil {
		return err
	}

	return nil
}

func (srv *AuthService) Register(data *RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	parsedDate, err := time.Parse(time.DateOnly, data.DateOfBirth)
	if err != nil {
		return err
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

	_, err = srv.UserRepository.Create(newUser)
	if err != nil {
		return err
	}

	return nil
}

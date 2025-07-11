package user

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("User not found.")
	ErrNoPermission = errors.New("You have no permissions to do the operation.")
)

type UserService struct {
	*UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (srv *UserService) Get(userID uint) (*User, error) {
	user, err := srv.UserRepository.Get(userID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *UserService) Update(user *User, authUserID uint) (*User, error) {
	return nil, nil
}

func (srv *UserService) Delete(userID, authUserID uint) error {
	user, err := srv.UserRepository.Get(userID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	if userID != authUserID {
		return ErrNoPermission
	}

	err = srv.UserRepository.Delete(userID)

	return err
}

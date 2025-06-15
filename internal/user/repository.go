package user

import (
	"restapi-sportshop/pkg/db"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Get(userID uint) (*User, error) {
	var user User

	result := repo.Database.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepository) GetByUsername(username string) (*User, error) {
	var user User

	result := repo.Database.DB.First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

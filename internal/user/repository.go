package user

import (
	"restapi-shop/pkg/db"
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

func (repo *UserRepository) GetIDByUsername(username string) (uint, error) {
	var user User

	result := repo.Database.DB.Select("id").First(&user, "username = ?", username)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) Delete(userID uint) error {
	result := repo.Database.DB.Delete(&User{}, userID)
	return result.Error
}

package item

import (
	"restapi-sportshop/pkg/db"
)

type ItemRepository struct {
	Database *db.Db
}

func NewItemRepository(database *db.Db) *ItemRepository {
	return &ItemRepository{
		Database: database,
	}
}

func (repo *ItemRepository) Get(itemID int) (*Item, error) {
	var item Item

	result := repo.Database.DB.First(&item, "id = ?", itemID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (repo *ItemRepository) Create(item *Item) (*Item, error) {
	result := repo.Database.DB.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

func (repo *ItemRepository) Update(item *Item) (*Item, error) {
	return nil, nil
}

func (repo *ItemRepository) Delete(itemID int) error {
	result := repo.Database.DB.Delete(&Item{}, itemID)
	return result.Error
}

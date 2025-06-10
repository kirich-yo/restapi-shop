package item

import (
	"restapi-sportshop/pkg/db"

	"gorm.io/gorm/clause"
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
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(item)
	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

func (repo *ItemRepository) Delete(itemID int) error {
	result := repo.Database.DB.Delete(&Item{}, itemID)
	return result.Error
}

func (repo *ItemRepository) Count(n *int64) error {
	result := repo.Database.DB.Table("items").Where("deleted_at IS NULL").Count(n)
	return result.Error
}

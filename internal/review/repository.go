package review

import (
	"restapi-sportshop/pkg/db"
)

type ReviewRepository struct {
	Database *db.Db
}

func NewReviewRepository(database *db.Db) *ReviewRepository {
	return &ReviewRepository{
		Database: database,
	}
}

func (repo *ReviewRepository) Get(reviewID uint) (*Review, error) {
	var res Review

	result := repo.Database.DB.First(&res, "id = ?", reviewID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &res, nil
}

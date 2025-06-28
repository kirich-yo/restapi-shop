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

func (repo *ReviewRepository) Create(review *Review) (*Review, error) {
	result := repo.Database.DB.Create(review)

	if result.Error != nil {
		return nil, result.Error
	}

	return review, nil
}

func (repo *ReviewRepository) Update(review *Review) (*Review, error) {
	return nil, nil
}

func (repo *ReviewRepository) Delete(reviewID uint) error {
	result := repo.Database.DB.Delete(&Review{}, reviewID)
	return result.Error
}

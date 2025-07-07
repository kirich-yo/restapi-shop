package review

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("Review not found.")
	ErrFKViolated = errors.New("Item not found.")
	ErrNoPermission = errors.New("You have no permissions to do the operation.")
)

type ReviewService struct {
	*ReviewRepository
}

func NewReviewService(reviewRepository *ReviewRepository) *ReviewService {
	return &ReviewService{
		ReviewRepository: reviewRepository,
	}
}

func (srv *ReviewService) Get(reviewID uint) (*Review, error) {
	review, err := srv.ReviewRepository.Get(reviewID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}

	return review, err
}

func (srv *ReviewService) Create(review *Review) (*Review, error) {
	review, err := srv.ReviewRepository.Create(review)

	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return nil, ErrFKViolated
	}

	return review, err
}

func (srv *ReviewService) Update(review *Review, authUserID uint) (*Review, error) {
	_, err := srv.ReviewRepository.Get(review.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if review.UserID != authUserID {
		return nil, ErrNoPermission
	}
	if err != nil {
		return nil, err
	}

	review, err = srv.ReviewRepository.Update(review)

	return review, err
}

func (srv *ReviewService) Delete(reviewID, authUserID uint) error {
	review, err := srv.ReviewRepository.Get(reviewID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	if review.UserID != authUserID {
		return ErrNoPermission
	}
	if err != nil {
		return err
	}

	err = srv.ReviewRepository.Delete(reviewID)

	return err
}

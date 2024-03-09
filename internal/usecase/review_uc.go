package usecase

import (
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
)

type ReviewUCItf interface {
	GetAll(page int) ([]*entity.Review, error)
	GetByRating(rating float64, page int) ([]*entity.Review, error)
}

type ReviewUC struct {
	repo repository.ReviewRepoItf
}

func (r ReviewUC) GetAll(page int) ([]*entity.Review, error) {
	limit := 5
	offset := (page - 1) * limit

	reviews, err := r.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return reviews, nil

}

func (r ReviewUC) GetByRating(rating float64, page int) ([]*entity.Review, error) {
	limit := 5
	offset := (page - 1) * limit

	reviews, err := r.repo.GetByRating(rating, limit, offset)
	if err != nil {
		return nil, err
	}

	return reviews, nil

}

func NewReviewUC(repo repository.ReviewRepoItf) ReviewUCItf {
	return &ReviewUC{repo: repo}
}

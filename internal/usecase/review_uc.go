package usecase

import (
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
	"projectIntern/model"
)

type ReviewUCItf interface {
	Create(req model.ReviewRequest) (*entity.Review, error)
	GetAll(page int) ([]*entity.Review, error)
}

type ReviewUC struct {
	repo repository.ReviewRepoItf
}

func (r ReviewUC) Create(req model.ReviewRequest) (*entity.Review, error) {
	review := &entity.Review{
		UserID:    req.UserId,
		ServiceID: req.ServiceId,
		Rating:    req.Rating,
		Review:    req.Review,
	}

	err := r.repo.Create(review)
	if err != nil {
		return nil, err
	}

	return review, nil
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

func NewReviewUC(repo repository.ReviewRepoItf) ReviewUCItf {
	return &ReviewUC{repo: repo}
}

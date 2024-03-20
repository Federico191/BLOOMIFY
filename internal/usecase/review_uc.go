package usecase

import (
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
)

type TreatmentReviewUCItf interface {
	GetAll(page int) ([]*entity.TreatmentReview, error)
}

type ReviewUC struct {
	repo repository.TreatmentReviewRepoItf
}

func NewReviewUC(repo repository.TreatmentReviewRepoItf) TreatmentReviewUCItf {
	return &ReviewUC{repo: repo}
}

func (r ReviewUC) GetAll(page int) ([]*entity.TreatmentReview, error) {
	limit := 5
	offset := (page - 1) * limit

	reviews, err := r.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return reviews, nil

}

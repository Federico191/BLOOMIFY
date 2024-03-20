package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/pkg/customerrors"
	"time"
)

type TreatmentReviewRepoItf interface {
	GetAll(limit, offset int) ([]*entity.TreatmentReview, error)
	GetById(id uint) (*entity.TreatmentReview, error)
}

type TreatmentReviewRepo struct {
	db *gorm.DB
}

func NewTreatmentReviewRepo(db *gorm.DB) TreatmentReviewRepoItf {
	return &TreatmentReviewRepo{db: db}
}

func (r TreatmentReviewRepo) GetAll(limit, offset int) ([]*entity.TreatmentReview, error) {
	var reviews []*entity.TreatmentReview
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := r.db.WithContext(ctx).Debug().Limit(limit).Offset(offset).Find(&reviews).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return reviews, nil
}

func (r TreatmentReviewRepo) GetById(id uint) (*entity.TreatmentReview, error) {
	var review *entity.TreatmentReview

	err := r.db.Debug().Where("id = ?", id).First(&review).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return review, nil
}

package repository

import (
	"errors"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/pkg/customerrors"
)

type DoctorReviewRepoItf interface {
	GetAll(limit, offset int) ([]*entity.DoctorReview, error)
	GetById(id uint) (*entity.DoctorReview, error)
}

type DoctorReviewRepo struct {
	db *gorm.DB
}

func NewDoctorReviewRepo(db *gorm.DB) DoctorReviewRepoItf {
	return &DoctorReviewRepo{db: db}
}

func (d DoctorReviewRepo) GetAll(limit, offset int) ([]*entity.DoctorReview, error) {
	var reviews []*entity.DoctorReview

	err := d.db.Debug().Limit(limit).Offset(offset).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (d DoctorReviewRepo) GetById(id uint) (*entity.DoctorReview, error) {
	var review *entity.DoctorReview

	err := d.db.Debug().Where("id = ?", id).First(&review).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return review, nil
}

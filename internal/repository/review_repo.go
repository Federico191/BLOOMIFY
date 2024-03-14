package repository

import (
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type ReviewRepoItf interface {
	Create(review *entity.Review) error
	GetAll(limit, offset int) ([]*entity.Review, error)
	GetById(id uint) (*entity.Review, error)
}

type ReviewRepo struct {
	db *gorm.DB
}

func NewReviewRepo(db *gorm.DB) ReviewRepoItf {
	return &ReviewRepo{db: db}
}

func (r ReviewRepo) Create(review *entity.Review) error {
	err := r.db.Debug().Create(&review).Error
	if err != nil {
		return err
	}

	return nil
}

func (r ReviewRepo) GetAll(limit, offset int) ([]*entity.Review, error) {
	var reviews []*entity.Review

	err := r.db.Debug().Limit(limit).Offset(offset).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r ReviewRepo) GetById(id uint) (*entity.Review, error) {
	var review *entity.Review

	err := r.db.Debug().Where("id = ?", id).First(&review).Error
	if err != nil {
		return nil, err
	}

	return review, nil
}

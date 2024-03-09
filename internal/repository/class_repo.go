package repository

import (
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"time"
)

type ClassRepoItf interface {
	Create(review *entity.Class) error
	GetAll(limit, offset int) ([]*entity.Class, error)
	GetByDay(day time.Time, limit, offset int) ([]*entity.Class, error)
	GetByRating(rating float64, limit, offset int) ([]*entity.Class, error)
}

type ClassRepo struct {
	db *gorm.DB
}

func (c ClassRepo) GetByDay(day time.Time, limit, offset int) ([]*entity.Class, error) {
	var classes []*entity.Class

	err := c.db.Debug().Where("day = ?", day).Limit(limit).Offset(offset).Find(&classes).Error
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func NewClassRepo(db *gorm.DB) ClassRepoItf {
	return &ClassRepo{db: db}
}

func (c ClassRepo) Create(class *entity.Class) error {
	err := c.db.Debug().Create(&class).Error
	if err != nil {
		return err
	}

	return nil
}

func (c ClassRepo) GetAll(limit, offset int) ([]*entity.Class, error) {
	var classes []*entity.Class

	err := c.db.Debug().Limit(limit).Offset(offset).Find(&classes).Error
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (c ClassRepo) GetByRating(rating float64, limit, offset int) ([]*entity.Class, error) {
	var classes []*entity.Class

	err := c.db.Debug().Where("rating = ?", rating).Limit(limit).Offset(offset).Find(&classes).Error
	if err != nil {
		return nil, err
	}

	return classes, nil
}

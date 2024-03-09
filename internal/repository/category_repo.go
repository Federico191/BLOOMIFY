package repository

import (
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type CategoryRepoItf interface {
	Create(category *entity.Category) error
	GetAll(limit, offset int) ([]*entity.Category, error)
}

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepoItf {
	return &CategoryRepo{db: db}
}

func (c CategoryRepo) Create(category *entity.Category) error {
	err := c.db.Debug().Create(&category).Error
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepo) GetAll(limit, offset int) ([]*entity.Category, error) {
	var categories []*entity.Category

	err := c.db.Debug().Limit(limit).Offset(offset).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

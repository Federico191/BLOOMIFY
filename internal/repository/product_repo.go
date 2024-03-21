package repository

import (
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type ProductRepoItf interface {
	GetByProblem(problemId uint) ([]*entity.Product, error)
}

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepoItf {
	return &ProductRepo{db: db}
}

func (p ProductRepo) GetByProblem(problemId uint) ([]*entity.Product, error) {
	var products []*entity.Product

	err := p.db.Debug().Preload("Problem").Where("problem_id = ?", problemId).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

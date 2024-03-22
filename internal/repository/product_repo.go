package repository

import (
	"errors"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/pkg/customerrors"
)

type ProductRepoItf interface {
	GetByProblem(problemId uint) ([]*entity.Product, error)
	GetByTopRate() ([]*entity.Product, error)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return products, nil
}

func (p ProductRepo) GetByTopRate() ([]*entity.Product, error) {
	var products []*entity.Product

	err := p.db.Debug().Preload("Problem").Order("rating desc").Limit(4).Find(&products).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return products, nil
}

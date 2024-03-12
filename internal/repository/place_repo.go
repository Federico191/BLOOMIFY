package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/model"
)

type PlaceRepoItf interface {
	GetById(id uuid.UUID) (*entity.Place, error)
	GetAllBeautyClinic(filter model.FilterParam, limit, offset int) ([]*entity.Place, error)
	GetAllSalon(filter model.FilterParam, limit, offset int) ([]*entity.Place, error)
	GetAllSpaMassage(filter model.FilterParam, limit, offset int) ([]*entity.Place, error)
	GetAllFitnessCenter(filter model.FilterParam, limit, offset int) ([]*entity.Place, error)
}

type PlaceRepo struct {
	db *gorm.DB
}

func (p PlaceRepo) GetById(id uuid.UUID) (*entity.Place, error) {
	var place *entity.Place

	err := p.db.Debug().Where("id = ?", id).First(&place).Error
	if err != nil {
		return nil, err
	}

	return place, nil
}

func (p PlaceRepo) GetAllBeautyClinic(filter model.FilterParam, limit, offset int) ([]*entity.Place, error) {
	var beautyClinics []*entity.Place

	if filter.MaxPrice == 0 {
		filter.MaxPrice = 1000000000
	}
	err := p.db.Debug().
		Preload("Category").
		Preload("Service", p.db.Debug().Where("price >= ? && price <= ?", filter.MinPrice, filter.MaxPrice)).
		Preload("Review", p.db.Debug().Where("rating = ?", filter.Rating)).
		Where("category_id = ? AND city LIKE ?", 1, "%"+filter.City+"%").
		Limit(limit).Offset(offset).Find(&beautyClinics).Error
	if err != nil {
		return nil, err
	}

	return beautyClinics, err
}

func (p PlaceRepo) GetAllSalon(filter model.FilterParam, limit, offset int) ([]*entity.Place, error) {
	var salons []*entity.Place

	if filter.MaxPrice == 0 {
		filter.MaxPrice = 1000000000
	}
	err := p.db.Debug().
		Preload("Category").
		Preload("Service", p.db.Debug().Where("price >= ? && price <= ?", filter.MinPrice, filter.MaxPrice)).
		Preload("Review", p.db.Debug().Where("rating = ?", filter.Rating)).
		Where("category_id = ? AND city LIKE ?", 3, "%"+filter.City+"%").
		Limit(limit).Offset(offset).Find(&salons).Error
	if err != nil {
		return nil, err
	}

	return salons, err
}

func (p PlaceRepo) GetAllSpaMassage(filter model.FilterParam, limit, offset int) ([]*entity.Place, error) {
	var spaMassages []*entity.Place

	if filter.MaxPrice == 0 {
		filter.MaxPrice = 1000000000
	}
	err := p.db.Debug().
		Preload("Category").
		Preload("Service", p.db.Debug().Where("price >= ? && price <= ?", filter.MinPrice, filter.MaxPrice)).
		Preload("Review", p.db.Debug().Where("rating = ?", filter.Rating)).
		Where("category_id = ? AND city LIKE ?", 2, "%"+filter.City+"%").
		Limit(limit).Offset(offset).Find(&spaMassages).Error
	if err != nil {
		return nil, err
	}

	return spaMassages, err
}

func (p PlaceRepo) GetAllFitnessCenter(filter model.FilterParam, limit, offset int) ([]*entity.Place, error) {
	var fitnessCenters []*entity.Place

	if filter.MaxPrice == 0 {
		filter.MaxPrice = 1000000000
	}
	err := p.db.Debug().
		Preload("Category").
		Preload("Service", p.db.Debug().Where("price >= ? && price <= ?", filter.MinPrice, filter.MaxPrice)).
		Preload("Review", p.db.Debug().Where("rating = ?", filter.Rating)).
		Where("category_id = ? AND city LIKE ?", 4, "%"+filter.City+"%").
		Limit(limit).Offset(offset).Find(&fitnessCenters).Error
	if err != nil {
		return nil, err
	}

	return fitnessCenters, err
}

func NewPlace(db *gorm.DB) PlaceRepoItf {
	return &PlaceRepo{db: db}
}

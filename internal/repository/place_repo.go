package repository

import (
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type PlaceRepoItf interface {
	GetById(id uint) (*entity.Place, error)
}

type PlaceRepo struct {
	db *gorm.DB
}

func NewPlaceRepo(db *gorm.DB) PlaceRepoItf {
	return &PlaceRepo{db: db}
}

func (p PlaceRepo) GetById(id uint) (*entity.Place, error) {
	var place *entity.Place

	err := p.db.Debug().Where("id = ?", id).First(&place).Error
	if err != nil {
		return nil, err
	}

	return place, nil
}

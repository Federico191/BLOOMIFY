package repository

import (
	"errors"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/pkg/customerrors"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return place, nil
}

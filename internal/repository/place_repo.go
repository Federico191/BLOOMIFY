package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type PlaceRepoItf interface {
	Create(place *entity.Place) (*entity.Place, error)
	GetById(id uuid.UUID) (*entity.Place, error)
	GetByCity(city string, limit, offset int) ([]*entity.Place, error)
	GetByAll(limit, offset int) ([]*entity.Place, error)
	GetTreatment(id uuid.UUID) (*entity.Place, error)
	GetClass(id uuid.UUID) (*entity.Place, error)
	GetReview(id uuid.UUID) (*entity.Place, error)
}

type PlaceRepo struct {
	db *gorm.DB
}

func (p PlaceRepo) GetReview(id uuid.UUID) (*entity.Place, error) {
	var place *entity.Place

	err := p.db.Debug().Where("id = ?", id).Preload("reviews").First(&place).Error
	if err != nil {
		return nil, err
	}

	return place, nil
}

func (p PlaceRepo) GetTreatment(id uuid.UUID) (*entity.Place, error) {
	var place *entity.Place

	err := p.db.Debug().Where("id = ?", id).Preload("services").First(&place).Error
	if err != nil {
		return nil, err
	}

	return place, nil
}

func (p PlaceRepo) GetClass(id uuid.UUID) (*entity.Place, error) {
	var place *entity.Place

	err := p.db.Debug().Where("id = ?", id).Preload("classes").First(&place).Error
	if err != nil {
		return nil, err
	}

	return place, nil
}

func NewPlace(db *gorm.DB) PlaceRepoItf {
	return &PlaceRepo{db: db}
}

func (p PlaceRepo) GetById(id uuid.UUID) (*entity.Place, error) {
	var place *entity.Place

	err := p.db.Debug().Where("id = ?", id).First(id).Error
	if err != nil {
		return nil, err
	}

	return place, nil
}

func (p PlaceRepo) GetByCity(city string, limit, offset int) ([]*entity.Place, error) {
	var places []*entity.Place

	err := p.db.Debug().Where("city = ?", city).Limit(limit).Offset(offset).Find(&places).Error
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (p PlaceRepo) GetByAll(limit, offset int) ([]*entity.Place, error) {
	var places []*entity.Place

	err := p.db.Debug().Limit(limit).Offset(offset).Find(&places).Error
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (p PlaceRepo) Create(place *entity.Place) (*entity.Place, error) {
	//TODO implement me
	panic("implement me")
}

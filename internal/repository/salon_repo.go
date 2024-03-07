package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type SalonRepoItf interface {
	Create(salon *entity.Salon) error
	GetById(id uuid.UUID) (*entity.Salon, error)
	GetByAddress(city, address string) (*entity.Salon, error)
	Update(clinic *entity.Salon) error
	GetByCity(city string, limit, offset int) ([]*entity.Salon, error)
	GetAll(limit, offset int) ([]*entity.Salon, error)
}

type SalonRepo struct {
	db *gorm.DB
}

func NewSalonRepo(db *gorm.DB) SalonRepoItf {
	return &SalonRepo{db: db}
}

func (s SalonRepo) Create(salon *entity.Salon) error {
	err := s.db.Debug().Create(salon).Error
	if err != nil {
		return err
	}

	return nil
}

func (s SalonRepo) GetById(id uuid.UUID) (*entity.Salon, error) {
	var salon *entity.Salon

	err := s.db.Debug().Where("id = ?", id).First(&salon).Error
	if err != nil {
		return nil, err
	}

	return salon, nil
}

func (s SalonRepo) GetByAddress(city, address string) (*entity.Salon, error) {
	var salon *entity.Salon

	err := s.db.Debug().Where("city = ? && address = ?", city, address).First(&salon).Error
	if err != nil {
		return nil, err
	}

	return salon, nil
}

func (s SalonRepo) Update(clinic *entity.Salon) error {
	err := s.db.Debug().Where("id = ?", clinic.ID).Updates(clinic).Error
	if err != nil {
		return err
	}

	return nil
}

func (s SalonRepo) GetByCity(city string, limit, offset int) ([]*entity.Salon, error) {
	var salon []*entity.Salon

	err := s.db.Debug().Where("city = ?", city).Limit(limit).Offset(offset).Find(&salon).Error
	if err != nil {
		return nil, err
	}

	return salon, nil
}

func (s SalonRepo) GetAll(limit, offset int) ([]*entity.Salon, error) {
	var salon []*entity.Salon

	err := s.db.Debug().Limit(limit).Offset(offset).Find(&salon).Error
	if err != nil {
		return nil, err
	}

	return salon, nil
}

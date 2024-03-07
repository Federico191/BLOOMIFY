package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type FitnessCenterRepoItf interface {
	Create(fitness *entity.FitnessCenter) error
	GetById(id uuid.UUID) (*entity.FitnessCenter, error)
	GetByAddress(city, address string) (*entity.FitnessCenter, error)
	Update(fitness *entity.FitnessCenter) error
	GetByCity(city string, limit, offset int) ([]*entity.FitnessCenter, error)
	GetAll(limit, offset int) ([]*entity.FitnessCenter, error)
}

type FitnessCenterRepo struct {
	db *gorm.DB
}

func NewFitnessCenterRepo(db *gorm.DB) FitnessCenterRepoItf {
	return &FitnessCenterRepo{db: db}
}

func (f FitnessCenterRepo) Create(fitness *entity.FitnessCenter) error {
	err := f.db.Debug().Create(&fitness).Error
	if err != nil {
		return err
	}

	return nil
}

func (f FitnessCenterRepo) GetById(id uuid.UUID) (*entity.FitnessCenter, error) {
	var fitnessCenter *entity.FitnessCenter

	err := f.db.Debug().Where("id = ?", id).First(&fitnessCenter).Error
	if err != nil {
		return nil, err
	}

	return fitnessCenter, nil
}

func (f FitnessCenterRepo) GetByAddress(city, address string) (*entity.FitnessCenter, error) {
	var fitnessCenter *entity.FitnessCenter

	err := f.db.Debug().Where("city = ? && address = ?", city, address).First(&fitnessCenter).Error
	if err != nil {
		return nil, err
	}

	return fitnessCenter, nil
}

func (f FitnessCenterRepo) Update(fitness *entity.FitnessCenter) error {
	//TODO implement me
	panic("implement me")
}

func (f FitnessCenterRepo) GetByCity(city string, limit, offset int) ([]*entity.FitnessCenter, error) {
	var fitnessCenters []*entity.FitnessCenter

	err := f.db.Debug().Where("City = ?", city).Limit(limit).Offset(offset).Find(&fitnessCenters).Error
	if err != nil {

	}

	return fitnessCenters, nil
}

func (f FitnessCenterRepo) GetAll(limit, offset int) ([]*entity.FitnessCenter, error) {
	var fitnessCenters []*entity.FitnessCenter

	err := f.db.Debug().Limit(limit).Offset(offset).Find(&fitnessCenters).Error
	if err != nil {

	}

	return fitnessCenters, nil
}

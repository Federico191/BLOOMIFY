package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type SpaMassageRepoItf interface {
	Create(salon *entity.SpaMassage) error
	GetById(id uuid.UUID) (*entity.SpaMassage, error)
	GetByAddress(city, address string) (*entity.SpaMassage, error)
	Update(clinic *entity.SpaMassage) error
	GetByCity(city string, limit, offset int) ([]*entity.SpaMassage, error)
	GetAll(limit, offset int) ([]*entity.SpaMassage, error)
}

type SpaMassageRepo struct {
	db *gorm.DB
}

func NewSpaMassageRepo(db *gorm.DB) SpaMassageRepoItf {
	return SpaMassageRepo{db: db}
}

func (s SpaMassageRepo) Create(salon *entity.SpaMassage) error {
	err := s.db.Debug().Create(&salon).Error
	if err != nil {
		return err
	}

	return nil
}

func (s SpaMassageRepo) GetByName(name string) (*entity.SpaMassage, error) {
	var spaMassages *entity.SpaMassage

	err := s.db.Debug().Where("name = ?", name).First(&spaMassages).Error
	if err != nil {
		return nil, err
	}

	return spaMassages, nil
}

func (s SpaMassageRepo) GetById(id uuid.UUID) (*entity.SpaMassage, error) {
	var spaMassage *entity.SpaMassage

	err := s.db.Debug().Where("id = ?", id).First(&spaMassage).Error
	if err != nil {
		return nil, err
	}

	return spaMassage, nil
}

func (s SpaMassageRepo) GetByAddress(city, address string) (*entity.SpaMassage, error) {
	var spaMassages *entity.SpaMassage

	err := s.db.Debug().Where("city = ? && address = ?", city, address).First(&spaMassages).Error
	if err != nil {
		return nil, err
	}

	return spaMassages, nil
}

func (s SpaMassageRepo) Update(clinic *entity.SpaMassage) error {
	//TODO implement me
	panic("implement me")
}

func (s SpaMassageRepo) GetByCity(city string, limit, offset int) ([]*entity.SpaMassage, error) {
	var spaMassages []*entity.SpaMassage

	err := s.db.Debug().Where("city = ?", city).Limit(limit).Offset(offset).Find(&spaMassages).Error
	if err != nil {
		return nil, err
	}

	return spaMassages, nil
}

func (s SpaMassageRepo) GetAll(limit, offset int) ([]*entity.SpaMassage, error) {
	var spaMassages []*entity.SpaMassage

	err := s.db.Debug().Limit(limit).Offset(offset).Find(&spaMassages).Error
	if err != nil {
		return nil, err
	}

	return spaMassages, nil
}

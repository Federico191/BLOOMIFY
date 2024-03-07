package repository

import (
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type BeautyClinicRepoItf interface {
	Create(clinic *entity.BeautyClinic) error
	GetById(id string) (*entity.BeautyClinic, error)
	GetByAddress(city, address string) (*entity.BeautyClinic, error)
	Update(clinic *entity.BeautyClinic) error
	GetByCity(city string, limit, offset int) ([]*entity.BeautyClinic, error)
	GetAll(limit, offset int) ([]*entity.BeautyClinic, error)
}

type BeautyClinicRepo struct {
	db *gorm.DB
}

func NewBeautyClinicRepo(db *gorm.DB) BeautyClinicRepoItf {
	return &BeautyClinicRepo{db: db}
}

func (b BeautyClinicRepo) Create(clinic *entity.BeautyClinic) error {
	err := b.db.Debug().Create(&clinic).Error
	if err != nil {
		return err
	}

	return nil
}

func (b BeautyClinicRepo) GetAll(limit, offset int) ([]*entity.BeautyClinic, error) {
	var beautyClinics []*entity.BeautyClinic

	err := b.db.Debug().Limit(limit).Offset(offset).Find(&beautyClinics).Error
	if err != nil {

	}

	return beautyClinics, nil
}

func (b BeautyClinicRepo) GetById(id string) (*entity.BeautyClinic, error) {
	var beautyClinic *entity.BeautyClinic

	err := b.db.Debug().Where("id = ?", id).First(&beautyClinic).Error
	if err != nil {
		return nil, err
	}

	return beautyClinic, nil
}

func (b BeautyClinicRepo) GetByAddress(city, address string) (*entity.BeautyClinic, error) {
	var beautyClinic *entity.BeautyClinic

	err := b.db.Debug().Where("city = ? && address = ?", city, address).First(&beautyClinic).Error
	if err != nil {
		return nil, err
	}

	return beautyClinic, nil
}

func (b BeautyClinicRepo) Update(clinic *entity.BeautyClinic) error {
	panic("implement")
}

func (b BeautyClinicRepo) GetByCity(city string, limit, offset int) ([]*entity.BeautyClinic, error) {
	var beautyClinics []*entity.BeautyClinic

	err := b.db.Debug().Where("City = ?", city).Limit(limit).Offset(offset).Find(&beautyClinics).Error
	if err != nil {

	}

	return beautyClinics, nil
}

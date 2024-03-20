package repository

import (
	"errors"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/pkg/customerrors"
	"time"
)

type BookingTreatmentRepoItf interface {
	Create(booking *entity.BookingTreatment) error
	GetById(id string) (*entity.BookingTreatment, error)
	GetByTimeId(id uint, time time.Time) (*entity.BookingTreatment, error)
	GetByStatus(transactionId string) (*entity.BookingTreatment, error)
	Update(booking *entity.BookingTreatment) error
	Delete(id string) error
}

type BookingTreatmentRepo struct {
	db *gorm.DB
}

func NewBookingTreatmentRepo(db *gorm.DB) BookingTreatmentRepoItf {
	return &BookingTreatmentRepo{db: db}
}

func (b BookingTreatmentRepo) GetByTimeId(id uint, time time.Time) (*entity.BookingTreatment, error) {
	var data entity.BookingTreatment
	if err := b.db.Debug().Where("service_id = ?", id).Where("book_at = ?", time).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (b BookingTreatmentRepo) Create(booking *entity.BookingTreatment) error {
	err := b.db.Debug().Create(&booking).Error
	if err != nil {
		return err
	}

	return nil
}

func (b BookingTreatmentRepo) GetById(id string) (*entity.BookingTreatment, error) {
	var booking *entity.BookingTreatment

	err := b.db.Debug().Where("id = ?", id).First(&booking).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return booking, nil
}

func (b BookingTreatmentRepo) GetByStatus(transactionId string) (*entity.BookingTreatment, error) {
	var booking *entity.BookingTreatment

	err := b.db.Debug().Preload("User").
		Preload("Service").
		Preload("Service.Place").
		Where("transaction_id = ?", transactionId).First(&booking).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return booking, nil
}

func (b BookingTreatmentRepo) Update(booking *entity.BookingTreatment) error {
	err := b.db.Debug().Updates(&booking).Error
	if err != nil {
		return err
	}

	return nil
}

func (b BookingTreatmentRepo) Delete(id string) error {
	var booking *entity.BookingTreatment

	err := b.db.Debug().Where("id = ?", id).Delete(&booking).Error
	if err != nil {
		return err
	}

	return nil
}

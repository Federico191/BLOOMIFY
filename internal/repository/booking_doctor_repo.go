package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/pkg/customerrors"
	"time"
)

type BookingDoctorRepoItf interface {
	Create(booking *entity.BookingDoctor) error
	GetById(id string) (*entity.BookingDoctor, error)
	GetByTimeId(id uuid.UUID, time time.Time) (*entity.BookingDoctor, error)
	GetByStatus(transactionId string) (*entity.BookingDoctor, error)
	Update(booking *entity.BookingDoctor) error
	Delete(id string) error
}

type BookingDoctorRepo struct {
	db *gorm.DB
}

func NewBookingDoctorRepo(db *gorm.DB) BookingDoctorRepoItf {
	return &BookingDoctorRepo{db: db}
}

func (b BookingDoctorRepo) Create(booking *entity.BookingDoctor) error {
	err := b.db.Debug().Create(&booking).Error
	if err != nil {
		return err
	}

	return nil
}

func (b BookingDoctorRepo) GetByTimeId(id uuid.UUID, time time.Time) (*entity.BookingDoctor, error) {
	var bookingDoctor *entity.BookingDoctor

	err := b.db.Debug().Where("doctor_id = ?", id).Where("book_at = ?", time).First(&bookingDoctor).Error
	if err != nil {
		return nil, err
	}
	return bookingDoctor, nil
}

func (b BookingDoctorRepo) GetByStatus(transactionId string) (*entity.BookingDoctor, error) {
	var BookingDoctor *entity.BookingDoctor

	err := b.db.Debug().Where("transaction_id = ?", transactionId).First(&BookingDoctor).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return BookingDoctor, nil
}

func (b BookingDoctorRepo) GetById(id string) (*entity.BookingDoctor, error) {
	var booking *entity.BookingDoctor
	err := b.db.Debug().Where("id = ?", id).First(&booking).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return booking, nil
}

func (b BookingDoctorRepo) Update(booking *entity.BookingDoctor) error {
	err := b.db.Debug().Updates(&booking).Error
	if err != nil {
		return err
	}

	return nil
}

func (b BookingDoctorRepo) Delete(id string) error {
	var BookingDoctor *entity.BookingDoctor

	err := b.db.Debug().Where("id = ?", id).Delete(&BookingDoctor).Error
	if err != nil {
		return err
	}

	return nil
}

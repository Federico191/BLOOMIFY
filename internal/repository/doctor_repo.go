package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/model"
	"projectIntern/pkg/customerrors"
)

type DoctorRepoItf interface {
	GetById(id uuid.UUID) (*entity.Doctor, error)
	GetAll(filter model.FilterParam, limit, offset int) ([]*entity.Doctor, error)
}

type DoctorRepo struct {
	db *gorm.DB
}

func NewDoctorRepo(db *gorm.DB) DoctorRepoItf {
	return &DoctorRepo{db: db}
}

func (d DoctorRepo) GetById(id uuid.UUID) (*entity.Doctor, error) {
	var doctor *entity.Doctor

	err := d.db.Debug().Preload("Profession").
		Preload("Reviews").
		Preload("Reviews.User").
		Where("id = ?", id).First(&doctor).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	// Menghitung rata-rata rating dari ulasan
	var totalRating int
	for _, review := range doctor.Reviews {
		totalRating += review.Rating
	}
	if len(doctor.Reviews) > 0 {
		doctor.Rating = float64(totalRating) / float64(len(doctor.Reviews))
	} else {
		doctor.Rating = 0
	}

	return doctor, nil
}

func (d DoctorRepo) GetAll(filter model.FilterParam, limit, offset int) ([]*entity.Doctor, error) {
	var doctors []*entity.Doctor

	query := d.db.Debug().Preload("Profession").
		Where("city LIKE ?", "%"+filter.City+"%").
		Preload("Reviews").
		Joins("JOIN doctor_reviews ON doctor_reviews.doctor_id = doctors.id").
		Group("doctors.id").
		Select("doctors.*", "COALESCE(AVG(doctor_reviews.rating), 0) as avg_rating").
		Limit(limit).Offset(offset)

	if filter.Price == "lowest" {
		query = query.Order("price")
	} else if filter.Price == "highest" {
		query = query.Order("price desc")
	}

	if filter.Rating == "lowest" {
		query = query.Order("avg_rating")
	} else if filter.Rating == "highest" {
		query = query.Order("avg_rating desc")
	}

	if filter.Price == "" && filter.Rating == "" && filter.City == "" {
		query = query.Order("avg_rating desc")
	}

	err := query.Find(&doctors).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	for _, data := range doctors {
		doctorID := data.ID
		var avg float64
		d.db.Model(&entity.DoctorReview{}).Where("doctor_id = ?", doctorID).Select("AVG(rating) as avg_rating").
			Find(&avg)
		data.Rating = avg
	}
	return doctors, nil

}

package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/model"
	"projectIntern/pkg/customerrors"
	"time"
)

type ServiceRepoItf interface {
	GetById(id uint) (*entity.Service, error)
	GetAllBeautyClinic(filter model.FilterParam, limit, offset int) ([]*entity.Service, error)
	GetAllSpaMassage(filter model.FilterParam, limit, offset int) ([]*entity.Service, error)
	GetAllSalon(filter model.FilterParam, limit, offset int) ([]*entity.Service, error)
	GetAllFitnessCenter(filter model.FilterParam, limit, offset int) ([]*entity.Service, error)
	GetByTopRate() ([]*entity.Service, error)
	GetByProblem(problemId uint) ([]*entity.Service, error)
}

type ServiceRepo struct {
	db *gorm.DB
}

func NewServiceRepo(db *gorm.DB) ServiceRepoItf {
	return &ServiceRepo{db: db}
}

func (s ServiceRepo) GetAll() ([]*entity.Service, error) {
	var service []*entity.Service

	err := s.db.Debug().
		Preload("Place").Find(&service).Error

	if err != nil {
		return nil, err
	}

	return service, nil
}

func (s ServiceRepo) GetAllBeautyClinic(filter model.FilterParam, limit, offset int) ([]*entity.Service, error) {
	var services []*entity.Service
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	query := s.db.WithContext(ctx).Debug().
		Preload("Place").
		Joins("JOIN places ON places.id = services.place_id").
		Where("places.category_id = ? AND places.city LIKE ?", 1, "%"+filter.City+"%").
		Preload("Problem").
		Preload("Reviews").
		Joins("JOIN treatment_reviews ON treatment_reviews.service_id = services.id").
		Group("services.id").
		Select("services.*", "COALESCE(AVG(treatment_reviews.rating), 0) as avg_rating").
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

	err := query.Find(&services).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	for _, data := range services {
		serviceID := data.ID
		var avg float64
		s.db.Model(&entity.TreatmentReview{}).Where("service_id = ?", serviceID).Select("AVG(rating) as avg_rating").
			Find(&avg)
		data.AvgRating = avg
	}
	return services, nil
}

func (s ServiceRepo) GetAllSpaMassage(filter model.FilterParam, limit, offset int) ([]*entity.Service, error) {
	var services []*entity.Service
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	query := s.db.WithContext(ctx).Debug().
		Preload("Place").
		Joins("JOIN places ON places.id = services.place_id").
		Where("places.category_id = ? AND places.city LIKE ?", 2, "%"+filter.City+"%").
		Preload("Problem").
		Preload("Reviews").
		Joins("JOIN treatment_reviews ON reviews.service_id = services.id").
		Group("services.id").
		Select("services.*", "COALESCE(AVG(treatment_reviews.rating), 0) as avg_rating").
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

	s.db.Joins("JOIN treatment_reviews ON treatment_reviews.service_id = services.id")

	err := query.Find(&services).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	for _, data := range services {
		serviceID := data.ID
		var avg float64
		s.db.Model(&entity.TreatmentReview{}).Where("service_id = ?", serviceID).Select("AVG(rating) as avg_rating").
			Find(&avg)
		data.AvgRating = avg
	}
	return services, nil
}

func (s ServiceRepo) GetById(id uint) (*entity.Service, error) {
	var service *entity.Service
	var avg float64

	query := s.db.Debug().Preload("Place").
		Where("services.id = ?", id).
		Preload("Reviews")

	query.Preload("Reviews.User")

	err := query.Find(&service).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	s.db.Debug().Model(entity.TreatmentReview{}).Where("service_id = ?", id).Select("AVG(rating) as avg_rating").
		Find(&avg)
	service.AvgRating = avg

	return service, nil
}

func (s ServiceRepo) GetAllSalon(filter model.FilterParam, limit, offset int) ([]*entity.Service, error) {
	var services []*entity.Service
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	query := s.db.WithContext(ctx).Debug().
		Preload("Place").
		Joins("JOIN places ON places.id = services.place_id").
		Where("places.category_id = ? AND places.city LIKE ?", 3, "%"+filter.City+"%").
		Joins("JOIN treatment_reviews ON treatment_reviews.service_id = services.id").
		Group("services.id").
		Select("services.*", "COALESCE(AVG(treatment_reviews.rating), 0) as avg_rating").
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

	s.db.Joins("JOIN treatment_reviews ON treatment_reviews.service_id = services.id")

	err := query.Find(&services).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	for _, data := range services {
		serviceID := data.ID
		var avg float64
		s.db.Model(&entity.TreatmentReview{}).Where("service_id = ?", serviceID).Select("AVG(rating) as avg_rating").
			Find(&avg)
		data.AvgRating = avg
	}
	return services, nil
}

func (s ServiceRepo) GetAllFitnessCenter(filter model.FilterParam, limit, offset int) ([]*entity.Service, error) {
	var services []*entity.Service
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	query := s.db.WithContext(ctx).Debug().
		Preload("Place").
		Joins("JOIN places ON places.id = services.place_id").
		Where("places.category_id = ? AND places.city LIKE ?", 4, "%"+filter.City+"%").
		Preload("Problem").
		Preload("Reviews").
		Joins("JOIN treatment_reviews ON treatment_reviews.service_id = services.id").
		Group("services.id").
		Select("services.*", "COALESCE(AVG(treatment_reviews.rating), 0) as avg_rating").
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

	s.db.Joins("JOIN treatment_reviews ON treatment_reviews.service_id = services.id")

	err := query.Find(&services).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	for _, data := range services {
		serviceID := data.ID
		var avg float64
		s.db.Model(&entity.TreatmentReview{}).Where("service_id = ?", serviceID).Select("AVG(rating) as avg_rating").
			Find(&avg)
		data.AvgRating = avg
	}
	return services, nil
}

func (s ServiceRepo) GetByTopRate() ([]*entity.Service, error) {
	var services []*entity.Service

	err := s.db.Table("services").
		Preload("Place").
		Select("services.*, AVG(treatment_reviews.rating) AS avg_rating").
		Joins("LEFT JOIN treatment_reviews ON services.id = treatment_reviews.service_id").
		Group("services.id").
		Order("avg_rating DESC").
		Limit(4).
		Find(&services).Error
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (s ServiceRepo) GetByProblem(problemId uint) ([]*entity.Service, error) {
	var services []*entity.Service
	err := s.db.Table("services").Preload("Place").Preload("Problem").Where("problem_id = ?", problemId).
		Limit(4).
		Find(&services).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}
	return services, nil
}

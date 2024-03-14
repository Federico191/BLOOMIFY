package repository

import (
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/model"
)

type ServiceRepoItf interface {
	Create(review entity.Review) error
	GetById(id uint) (*entity.Service, error)
	GetAllBeautyClinic(filter model.FilterParam, limit, offset int) ([]*entity.Service, error)
	GetAllSpaMassage(filter model.FilterParam, limit, offset int) ([]*entity.Service, error)
	GetAllSalon(filter model.FilterParam, limit, offset int) ([]*entity.Service, error)
	GetAllFitnessCenter(filter model.FilterParam, limit, offset int) ([]*entity.Service, error)
}

type ServiceRepo struct {
	db *gorm.DB
}

func NewServiceRepo(db *gorm.DB) ServiceRepoItf {
	return &ServiceRepo{db: db}
}

func (s ServiceRepo) Create(review entity.Review) error {
	err := s.db.Debug().Create(&review).Error
	if err != nil {
		return err
	}

	return nil
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
	query := s.db.Debug().
		Preload("Place").
		Joins("JOIN places ON places.id = services.place_id").
		Where("places.category_id = ? AND places.city LIKE ?", 1, "%"+filter.City+"%").
		Preload("Problem").
		Preload("Reviews").
		Joins("JOIN reviews ON reviews.service_id = services.id").
		Group("services.id").
		Select("services.*", "COALESCE(AVG(reviews.rating), 0) as avg_rating").
		Limit(limit).Offset(offset)

	if filter.Price == "lower" {
		query = query.Order("price")
	} else if filter.Price == "highest" {
		query = query.Order("price desc")
	}

	if filter.Rating == "lower" {
		query = query.Order("avg_rating")
	} else if filter.Rating == "highest" {
		query = query.Order("avg_rating desc")
	}

	if filter.Price == "" && filter.Rating == "" && filter.City == "" {
		query = query.Order("avg_rating desc")
	}

	err := query.Find(&services).Error
	if err != nil {
		return nil, err
	}
	for _, data := range services {
		serviceID := data.ID
		var avg float64
		s.db.Model(&entity.Review{}).Where("service_id = ?", serviceID).Select("AVG(rating) as avg_rating").
			Find(&avg)
		data.AvgRating = avg
	}
	return services, nil
}

func (s ServiceRepo) GetAllSpaMassage(filter model.FilterParam, limit, offset int) ([]*entity.Service, error) {
	var services []*entity.Service

	query := s.db.Debug().
		Preload("Place").
		Joins("JOIN places ON places.id = services.place_id").
		Where("places.category_id = ? AND places.city LIKE ?", 2, "%"+filter.City+"%").
		Preload("Problem").
		Preload("Reviews").
		Joins("JOIN reviews ON reviews.service_id = services.id").
		Group("services.id").
		Select("services.*", "COALESCE(AVG(reviews.rating), 0) as avg_rating").
		Limit(limit).Offset(offset)

	if filter.Price == "lower" {
		query = query.Order("price")
	} else if filter.Price == "highest" {
		query = query.Order("price desc")
	}

	if filter.Rating == "lower" {
		query = query.Order("avg_rating")
	} else if filter.Rating == "highest" {
		query = query.Order("avg_rating desc")
	}

	if filter.Price == "" && filter.Rating == "" && filter.City == "" {
		query = query.Order("avg_rating desc")
	}

	s.db.Joins("JOIN reviews ON reviews.service_id = services.id")

	err := query.Find(&services).Error
	if err != nil {
		return nil, err
	}
	for _, data := range services {
		serviceID := data.ID
		var avg float64
		s.db.Model(&entity.Review{}).Where("service_id = ?", serviceID).Select("AVG(rating) as avg_rating").
			Find(&avg)
		data.AvgRating = avg
	}
	return services, nil
}

func (s ServiceRepo) GetById(id uint) (*entity.Service, error) {
	var service *entity.Service

	err := s.db.Debug().Where("id = ?", id).First(&service).Error
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (s ServiceRepo) GetAllSalon(filter model.FilterParam, limit, offset int) ([]*entity.Service, error) {
	var services []*entity.Service

	query := s.db.Debug().
		Preload("Place").
		Joins("JOIN places ON places.id = services.place_id").
		Where("places.category_id = ? AND places.city LIKE ?", 3, "%"+filter.City+"%").
		Preload("Problem").
		Preload("Reviews").
		Joins("JOIN reviews ON reviews.service_id = services.id").
		Group("services.id").
		Select("services.*", "COALESCE(AVG(reviews.rating), 0) as avg_rating").
		Limit(limit).Offset(offset)

	if filter.Price == "lower" {
		query = query.Order("price")
	} else if filter.Price == "highest" {
		query = query.Order("price desc")
	}

	if filter.Rating == "lower" {
		query = query.Order("avg_rating")
	} else if filter.Rating == "highest" {
		query = query.Order("avg_rating desc")
	}

	if filter.Price == "" && filter.Rating == "" && filter.City == "" {
		query = query.Order("avg_rating desc")
	}

	s.db.Joins("JOIN reviews ON reviews.service_id = services.id")

	err := query.Find(&services).Error
	if err != nil {
		return nil, err
	}
	for _, data := range services {
		serviceID := data.ID
		var avg float64
		s.db.Model(&entity.Review{}).Where("service_id = ?", serviceID).Select("AVG(rating) as avg_rating").
			Find(&avg)
		data.AvgRating = avg
	}
	return services, nil
}

func (s ServiceRepo) GetAllFitnessCenter(filter model.FilterParam, limit, offset int) ([]*entity.Service, error) {
	var services []*entity.Service

	query := s.db.Debug().
		Preload("Place").
		Joins("JOIN places ON places.id = services.place_id").
		Where("places.category_id = ? AND places.city LIKE ?", 4, "%"+filter.City+"%").
		Preload("Problem").
		Preload("Reviews").
		Joins("JOIN reviews ON reviews.service_id = services.id").
		Group("services.id").
		Select("services.*", "COALESCE(AVG(reviews.rating), 0) as avg_rating").
		Limit(limit).Offset(offset)

	if filter.Price == "lower" {
		query = query.Order("price")
	} else if filter.Price == "highest" {
		query = query.Order("price desc")
	}

	if filter.Rating == "lower" {
		query = query.Order("avg_rating")
	} else if filter.Rating == "highest" {
		query = query.Order("avg_rating desc")
	}

	if filter.Price == "" && filter.Rating == "" && filter.City == "" {
		query = query.Order("avg_rating desc")
	}

	s.db.Joins("JOIN reviews ON reviews.service_id = services.id")

	err := query.Find(&services).Error
	if err != nil {
		return nil, err
	}
	for _, data := range services {
		serviceID := data.ID
		var avg float64
		s.db.Model(&entity.Review{}).Where("service_id = ?", serviceID).Select("AVG(rating) as avg_rating").
			Find(&avg)
		data.AvgRating = avg
	}
	return services, nil
}

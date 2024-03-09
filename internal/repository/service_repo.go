package repository

import (
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type ServiceRepoItf interface {
	Create(review entity.Review) error
	GetAll(limit, offset int) ([]*entity.Service, error)
	GetByTreatment(treatment string, limit, offset int) ([]*entity.Service, error)
	GetByProblem(problem string, limit, offset int) ([]*entity.Service, error)
}

type ServiceRepo struct {
	db *gorm.DB
}

func (s ServiceRepo) Create(review entity.Review) error {
	err := s.db.Debug().Create(&review).Error
	if err != nil {
		return err
	}

	return nil
}

func (s ServiceRepo) GetAll(limit, offset int) ([]*entity.Service, error) {
	var services []*entity.Service

	err := s.db.Debug().Limit(limit).Offset(offset).Find(&services).Error
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (s ServiceRepo) GetByTreatment(treatment string, limit, offset int) ([]*entity.Service, error) {
	var services []*entity.Service

	err := s.db.Debug().Where("treatment = ?", treatment).Limit(limit).Offset(offset).Find(&services).Error
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (s ServiceRepo) GetByProblem(problem string, limit, offset int) ([]*entity.Service, error) {
	var services []*entity.Service

	err := s.db.Debug().Where("problem = ?", problem).Limit(limit).Offset(offset).Find(&services).Error
	if err != nil {
		return nil, err
	}

	return services, nil
}

func NewServiceRepo(db *gorm.DB) ServiceRepoItf {
	return &ServiceRepo{db: db}
}

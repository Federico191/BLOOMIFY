package usecase

import (
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
)

type ServiceItf interface {
	GetAll(page int) ([]*entity.Service, error)
	GetByTreatment(treatment string, page int) ([]*entity.Service, error)
	GetByProblem(problem string, page int) ([]*entity.Service, error)
}

type Service struct {
	repo repository.ServiceRepoItf
}

func NewService(repo repository.ServiceRepoItf) ServiceItf {
	return &Service{repo: repo}
}

func (s Service) GetAll(page int) ([]*entity.Service, error) {
	limit := 5
	offset := (page - 1) * limit

	services, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (s Service) GetByTreatment(treatment string, page int) ([]*entity.Service, error) {
	limit := 5
	offset := (page - 1) * limit

	services, err := s.repo.GetByTreatment(treatment, limit, offset)
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (s Service) GetByProblem(problem string, page int) ([]*entity.Service, error) {
	limit := 5
	offset := (page - 1) * limit

	services, err := s.repo.GetByTreatment(problem, limit, offset)
	if err != nil {
		return nil, err
	}

	return services, nil
}

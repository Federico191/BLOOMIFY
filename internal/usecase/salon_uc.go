package usecase

import (
	"errors"
	"github.com/google/uuid"
	"projectIntern/internal/entity"
	"projectIntern/internal/model"
	"projectIntern/internal/repository"
	"projectIntern/pkg/customerrors"
)

type SalonUCItf interface {
	Create(req model.SalonRequest) (*entity.Salon, error)
	Update(req model.SalonUpdate) (*entity.Salon, error)
	GetByCity(city string, page int) ([]*entity.Salon, error)
	GetAll(page int) ([]*entity.Salon, error)
}

type SalonUC struct {
	repo repository.SalonRepoItf
}

func NewSalonUC(repo repository.SalonRepoItf) SalonUCItf {
	return &SalonUC{repo: repo}
}

func (s SalonUC) Create(req model.SalonRequest) (*entity.Salon, error) {
	exist, err := s.repo.GetByAddress(req.City, req.Address)
	if err != nil {
		return nil, err
	}

	if exist != nil {
		return nil, customerrors.ErrRecordAlreadyExist
	}

	salon := &entity.Salon{
		ID:        uuid.New(),
		Name:      req.Name,
		Address:   req.Address,
		City:      req.City,
		Contact:   req.Contact,
		Hour:      req.Hour,
		PhotoLink: req.PhotoLink,
	}

	err = s.repo.Create(salon)
	if err != nil {
		return nil, errors.New("cannot create salon")
	}

	return salon, err
}

func (s SalonUC) Update(req model.SalonUpdate) (*entity.Salon, error) {
	salon, err := s.repo.GetById(req.ID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(salon)
	if err != nil {
		return nil, err
	}

	return salon, nil
}

func (s SalonUC) GetByCity(city string, page int) ([]*entity.Salon, error) {
	limit := 5
	offset := (page - 1) * limit

	salons, err := s.repo.GetByCity(city, limit, offset)
	if err != nil {
		return nil, err
	}

	return salons, nil
}

func (s SalonUC) GetAll(page int) ([]*entity.Salon, error) {
	limit := 5
	offset := (page - 1) * limit

	salons, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return salons, nil
}

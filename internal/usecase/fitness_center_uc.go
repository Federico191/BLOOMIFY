package usecase

import (
	"errors"
	"github.com/google/uuid"
	"projectIntern/internal/entity"
	"projectIntern/internal/model"
	"projectIntern/internal/repository"
	"projectIntern/pkg/customerrors"
)

type FitnessCenterUCItf interface {
	Create(req model.FitnessCenterRequest) (*entity.FitnessCenter, error)
	Update(req model.FitnessCenterUpdate) (*entity.FitnessCenter, error)
	GetByCity(city string, page int) ([]*entity.FitnessCenter, error)
	GetAll(page int) ([]*entity.FitnessCenter, error)
}

type FitnessCenterUC struct {
	repo repository.FitnessCenterRepoItf
}

func NewFitnessCenterUc(repo repository.FitnessCenterRepoItf) FitnessCenterUCItf {
	return &FitnessCenterUC{repo: repo}
}

func (f FitnessCenterUC) Create(req model.FitnessCenterRequest) (*entity.FitnessCenter, error) {
	exist, err := f.repo.GetByAddress(req.City, req.Address)
	if err != nil {
		return nil, err
	}

	if exist != nil {
		return nil, customerrors.ErrRecordAlreadyExist
	}

	fc := &entity.FitnessCenter{
		ID:        uuid.New(),
		Name:      req.Name,
		Address:   req.Address,
		City:      req.City,
		Contact:   req.Contact,
		Hour:      req.Hour,
		PhotoLink: req.PhotoLink,
	}

	err = f.repo.Create(fc)
	if err != nil {
		return nil, errors.New("cannot insert data")
	}

	return fc, nil
}

func (f FitnessCenterUC) GetAll(page int) ([]*entity.FitnessCenter, error) {
	limit := 5
	offset := (page - 1) * limit

	fitnessCenters, err := f.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return fitnessCenters, nil
}

func (f FitnessCenterUC) Update(req model.FitnessCenterUpdate) (*entity.FitnessCenter, error) {
	beautyClinic, err := f.repo.GetById(req.ID)
	if err != nil {
		return nil, err
	}

	err = f.repo.Update(beautyClinic)
	if err != nil {
		return nil, err
	}

	return beautyClinic, nil
}

func (f FitnessCenterUC) GetByCity(city string, page int) ([]*entity.FitnessCenter, error) {
	limit := 5
	offset := (page - 1) * limit

	fitnessCenters, err := f.repo.GetByCity(city, limit, offset)
	if err != nil {
		return nil, err
	}

	if fitnessCenters == nil {
		return nil, customerrors.ErrRecordNotFound
	}

	return fitnessCenters, nil
}

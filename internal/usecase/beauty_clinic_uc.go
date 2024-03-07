package usecase

import (
	"errors"
	"github.com/google/uuid"
	"projectIntern/internal/entity"
	"projectIntern/internal/model"
	"projectIntern/internal/repository"
	"projectIntern/pkg/customerrors"
)

type BeautyClinicUCItf interface {
	Create(req model.BeautyClinicRequest) (*entity.BeautyClinic, error)
	Update(req model.BeautyClinicUpdate) (*entity.BeautyClinic, error)
	GetByCity(city string, page int) ([]*entity.BeautyClinic, error)
	GetAll(page int) ([]*entity.BeautyClinic, error)
}

type BeautyClinicUC struct {
	repo repository.BeautyClinicRepoItf
}

func NewBeautyClinicUC(repo repository.BeautyClinicRepoItf) BeautyClinicUCItf {
	return &BeautyClinicUC{repo: repo}
}

func (b BeautyClinicUC) Create(req model.BeautyClinicRequest) (*entity.BeautyClinic, error) {
	exist, err := b.repo.GetByAddress(req.City, req.Address)
	if err != nil {
		return nil, err
	}

	if exist != nil {
		return nil, customerrors.ErrRecordAlreadyExist
	}

	bc := &entity.BeautyClinic{
		ID:        uuid.New(),
		Name:      req.Name,
		Address:   req.Address,
		City:      req.City,
		Contact:   req.Contact,
		Hour:      req.Hour,
		PhotoLink: req.PhotoLink,
	}

	err = b.repo.Create(bc)
	if err != nil {
		return nil, errors.New("cannot insert data")
	}

	return bc, nil
}

func (b BeautyClinicUC) GetByCity(city string, page int) ([]*entity.BeautyClinic, error) {
	limit := 5
	offset := (page - 1) * limit

	beautyClinics, err := b.repo.GetByCity(city, limit, offset)
	if err != nil {
		return nil, err
	}

	if beautyClinics == nil {
		return nil, customerrors.ErrRecordNotFound
	}

	return beautyClinics, nil
}

func (b BeautyClinicUC) GetAll(page int) ([]*entity.BeautyClinic, error) {
	limit := 5
	offset := (page - 1) * limit

	beautyClinics, err := b.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return beautyClinics, nil
}

func (b BeautyClinicUC) Update(req model.BeautyClinicUpdate) (*entity.BeautyClinic, error) {
	beautyClinic, err := b.repo.GetById(req.ID)
	if err != nil {
		return nil, err
	}

	err = b.repo.Update(beautyClinic)
	if err != nil {
		return nil, err
	}

	return beautyClinic, nil
}

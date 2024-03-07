package usecase

import (
	"github.com/google/uuid"
	"projectIntern/internal/entity"
	"projectIntern/internal/model"
	"projectIntern/internal/repository"
	"projectIntern/pkg/customerrors"
)

type SpaMassageUCItf interface {
	Create(req model.SpaMassageRequest) (*entity.SpaMassage, error)
	Update(req model.SpaMassageUpdate) (*entity.SpaMassage, error)
	GetByCity(city string, page int) ([]*entity.SpaMassage, error)
	GetAll(page int) ([]*entity.SpaMassage, error)
}

type SpaMassageUC struct {
	repo repository.SpaMassageRepoItf
}

func NewSpaMassageUC(repo repository.SpaMassageRepoItf) SpaMassageUCItf {
	return SpaMassageUC{repo: repo}
}

func (s SpaMassageUC) Create(req model.SpaMassageRequest) (*entity.SpaMassage, error) {
	exist, err := s.repo.GetByAddress(req.City, req.Address)
	if err != nil {
		return nil, err
	}

	if exist != nil {
		return nil, customerrors.ErrRecordAlreadyExist
	}

	spaMassage := &entity.SpaMassage{
		ID:        uuid.New(),
		Name:      req.Name,
		Address:   req.Address,
		City:      req.City,
		Contact:   req.Contact,
		Hour:      req.Hour,
		PhotoLink: req.PhotoLink,
	}

	err = s.repo.Create(spaMassage)
	if err != nil {
		return nil, err
	}

	return spaMassage, nil
}

func (s SpaMassageUC) Update(req model.SpaMassageUpdate) (*entity.SpaMassage, error) {
	spaMassage, err := s.repo.GetById(req.ID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(spaMassage)
	if err != nil {
		return nil, err
	}

	return spaMassage, nil
}

func (s SpaMassageUC) GetByCity(city string, page int) ([]*entity.SpaMassage, error) {
	limit := 5
	offset := (page - 1) * limit

	spaMassages, err := s.repo.GetByCity(city, limit, offset)
	if err != nil {
		return nil, err
	}

	return spaMassages, nil
}

func (s SpaMassageUC) GetAll(page int) ([]*entity.SpaMassage, error) {
	limit := 5
	offset := (page - 1) * limit

	spaMassages, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return spaMassages, nil
}

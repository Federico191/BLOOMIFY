package usecase

import (
	"github.com/google/uuid"
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
)

type PlaceUCItf interface {
	GetAll(page int) ([]*entity.Place, error)
	GetByCity(city string, page int) ([]*entity.Place, error)
	GetTreatment(id uuid.UUID) (*entity.Place, error)
	GetClasses(id uuid.UUID) (*entity.Place, error)
}

type PlaceUC struct {
	placeRepo repository.PlaceRepoItf
}

func (p PlaceUC) GetTreatment(id uuid.UUID) (*entity.Place, error) {
	return p.placeRepo.GetTreatment(id)
}

func (p PlaceUC) GetClasses(id uuid.UUID) (*entity.Place, error) {
	return p.placeRepo.GetClass(id)
}

func (p PlaceUC) GetAll(page int) ([]*entity.Place, error) {
	limit := 5
	offset := (page - 1) * limit

	places, err := p.placeRepo.GetByAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (p PlaceUC) GetByCity(city string, page int) ([]*entity.Place, error) {
	limit := 5
	offset := (page - 1) * limit

	places, err := p.placeRepo.GetByCity(city, limit, offset)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func NewPlaceUC(placeRepo repository.PlaceRepoItf) PlaceUCItf {
	return &PlaceUC{placeRepo: placeRepo}
}

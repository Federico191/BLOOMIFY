package usecase

import (
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
	"projectIntern/model"
)

type PlaceUCItf interface {
	GetAllBeautyClinic(filter model.FilterParam, page int) ([]*entity.Place, error)
	GetAllSalon(filter model.FilterParam, page int) ([]*entity.Place, error)
	GetAllSpaMassage(filter model.FilterParam, page int) ([]*entity.Place, error)
	GetAllFitnessCenter(filter model.FilterParam, page int) ([]*entity.Place, error)
}

type PlaceUC struct {
	placeRepo repository.PlaceRepoItf
}

func NewPlaceUC(placeRepo repository.PlaceRepoItf) PlaceUCItf {
	return &PlaceUC{placeRepo: placeRepo}
}

func (p PlaceUC) GetAllBeautyClinic(filter model.FilterParam, page int) ([]*entity.Place, error) {
	limit := 5
	offset := (page - 1) * limit

	beautyClinics, err := p.placeRepo.GetAllBeautyClinic(filter, limit, offset)
	if err != nil {
		return nil, err
	}

	return beautyClinics, nil
}

func (p PlaceUC) GetAllSalon(filter model.FilterParam, page int) ([]*entity.Place, error) {
	limit := 5
	offset := (page - 1) * limit

	salons, err := p.placeRepo.GetAllSalon(filter, limit, offset)
	if err != nil {
		return nil, err
	}

	return salons, nil
}

func (p PlaceUC) GetAllSpaMassage(filter model.FilterParam, page int) ([]*entity.Place, error) {
	limit := 5
	offset := (page - 1) * limit

	spaMassages, err := p.placeRepo.GetAllSpaMassage(filter, limit, offset)
	if err != nil {
		return nil, err
	}

	return spaMassages, nil
}

func (p PlaceUC) GetAllFitnessCenter(filter model.FilterParam, page int) ([]*entity.Place, error) {
	limit := 5
	offset := (page - 1) * limit

	fitnessCenters, err := p.placeRepo.GetAllFitnessCenter(filter, limit, offset)
	if err != nil {
		return nil, err
	}

	return fitnessCenters, nil
}

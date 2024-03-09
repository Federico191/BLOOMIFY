package usecase

import (
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
	"time"
)

type ClassUCItf interface {
	GetAll(page int) ([]*entity.Class, error)
	GetByDay(day time.Time, page int) ([]*entity.Class, error)
	GetByRating(rating float64, page int) ([]*entity.Class, error)
}

type ClassUC struct {
	repo repository.ClassRepoItf
}

func (c ClassUC) GetAll(page int) ([]*entity.Class, error) {
	limit := 5
	offset := (page - 1) * limit

	classes, err := c.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (c ClassUC) GetByDay(day time.Time, page int) ([]*entity.Class, error) {
	limit := 5
	offset := (page - 1) * limit

	classes, err := c.repo.GetByDay(day, limit, offset)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (c ClassUC) GetByRating(rating float64, page int) ([]*entity.Class, error) {
	limit := 5
	offset := (page - 1) * limit

	classes, err := c.repo.GetByRating(rating, limit, offset)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func NewClassUc(repo repository.ClassRepoItf) ClassUCItf {
	return &ClassUC{repo: repo}
}

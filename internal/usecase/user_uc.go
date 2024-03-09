package usecase

import (
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
)

type UserUCItf interface {
	GetById(id string) (*entity.User, error)
}

type UserUC struct {
	repo repository.UserRepoItf
}

func NewUseUC(repo repository.UserRepoItf) UserUCItf {
	return &UserUC{repo: repo}
}

func (u UserUC) GetById(id string) (*entity.User, error) {
	user, err := u.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

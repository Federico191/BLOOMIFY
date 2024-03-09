package repository

import (
	"errors"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
)

type UserRepoItf interface {
	Create(user *entity.User) error
	GetByEmail(email string) (*entity.User, error)
	GetById(id string) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoItf {
	return UserRepository{db: db}
}

func (u UserRepository) Create(user *entity.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) GetByEmail(email string) (*entity.User, error) {
	var user *entity.User

	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

func (u UserRepository) GetById(id string) (*entity.User, error) {
	var user *entity.User

	err := u.db.Debug().Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

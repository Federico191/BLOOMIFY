package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/model"
)

type UserRepoItf interface {
	Create(user *entity.User) error
	GetByEmail(email string) (*entity.User, error)
	GetById(id uuid.UUID) (*entity.User, error)
	GetByVerificationCode(code string) (*entity.User, error)
	Update(user *entity.User, req model.UserUpdate) error
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
		return nil, err
	}

	return user, nil
}

func (u UserRepository) GetByVerificationCode(code string) (*entity.User, error) {
	var user *entity.User

	err := u.db.Debug().Where("verification_code = ?", code).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserRepository) GetById(id uuid.UUID) (*entity.User, error) {
	var user *entity.User

	err := u.db.Debug().Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserRepository) Update(user *entity.User, req model.UserUpdate) error {
	err := u.db.Debug().Model(&user).Updates(req).Error
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Delete(user *entity.User) error {
	err := u.db.Debug().Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

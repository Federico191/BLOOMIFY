package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/model"
	"projectIntern/pkg/customerrors"
	"time"
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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) GetByEmail(email string) (*entity.User, error) {
	var user *entity.User
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return user, nil
}

func (u UserRepository) GetByVerificationCode(code string) (*entity.User, error) {
	var user *entity.User
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := u.db.WithContext(ctx).Debug().Where("verification_code = ?", code).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, err
	}

	return user, nil
}

func (u UserRepository) GetById(id uuid.UUID) (*entity.User, error) {
	var user *entity.User
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := u.db.WithContext(ctx).Debug().Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
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

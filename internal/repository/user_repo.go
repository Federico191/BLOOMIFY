package repository

import (
	"context"
	"gorm.io/gorm"
	"log"
	"projectIntern/internal/entity"
)

type UserRepoItf interface {
	Create(ctx context.Context, user *entity.User) error
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoItf {
	return UserRepository{db: db}
}

func (u UserRepository) Create(ctx context.Context, user *entity.User) error {
	err := u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user *entity.User

	err := u.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return user, nil
}

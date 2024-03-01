package repository

import "gorm.io/gorm"

type Repository struct {
	User UserRepoItf
}

func Init(db *gorm.DB) *Repository {
	return &Repository{User: NewUserRepo(db)}
}

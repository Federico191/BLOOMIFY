package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	User     UserRepoItf
	Service  ServiceRepoItf
	Review   ReviewRepoItf
	Category CategoryRepoItf
}

func Init(db *gorm.DB) *Repository {
	return &Repository{
		User:     NewUserRepo(db),
		Service:  NewServiceRepo(db),
		Review:   NewReviewRepo(db),
		Category: NewCategoryRepo(db),
	}
}

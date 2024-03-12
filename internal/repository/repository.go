package repository

import (
	"gorm.io/gorm"
	"projectIntern/internal/repository/category"
)

type Repository struct {
	User     UserRepoItf
	Place    PlaceRepoItf
	Service  ServiceRepoItf
	Class    ClassRepoItf
	Review   ReviewRepoItf
	Category category.CategoryRepoItf
}

func Init(db *gorm.DB) *Repository {
	return &Repository{
		User:    NewUserRepo(db),
		Place:   NewPlace(db),
		Service: NewServiceRepo(db),
		Class:   NewClassRepo(db),
		Review:  NewReviewRepo(db),
	}
}

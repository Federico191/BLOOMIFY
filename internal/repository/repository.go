package repository

import "gorm.io/gorm"

type Repository struct {
	User     UserRepoItf
	Place    PlaceRepoItf
	Service  ServiceRepoItf
	Class    ClassRepoItf
	Review   ReviewRepoItf
	Category CategoryRepoItf
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

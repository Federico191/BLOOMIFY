package usecase

import (
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
	"projectIntern/model"
)

type ServiceItf interface {
	GetById(id uint) (*model.ServiceDetailResponse, error)
	GetAllBeautyClinic(filter model.FilterParam, page int) ([]*model.ServiceResponse, error)
	GetAllSpaMassage(filter model.FilterParam, page int) ([]*model.ServiceResponse, error)
	GetAllSalon(filter model.FilterParam, page int) ([]*model.ServiceResponse, error)
	GetAllFitnessCenter(filter model.FilterParam, page int) ([]*model.ServiceResponse, error)
}

type Service struct {
	serviceRepo  repository.ServiceRepoItf
	categoryRepo repository.CategoryRepoItf
	reviewRepo   repository.ReviewRepoItf
	userRepo     repository.UserRepoItf
}

func NewService(repo repository.ServiceRepoItf, categoryRepo repository.CategoryRepoItf) ServiceItf {
	return &Service{serviceRepo: repo, categoryRepo: categoryRepo}
}

func (s Service) GetAllBeautyClinic(filter model.FilterParam, page int) ([]*model.ServiceResponse, error) {
	limit := 5
	offset := (page - 1) * limit

	beautyClinics, err := s.serviceRepo.GetAllBeautyClinic(filter, limit, offset)
	if err != nil {
		return nil, err
	}

	var services []*model.ServiceResponse
	for _, beautyClinic := range beautyClinics {
		category, err := s.categoryRepo.GetById(beautyClinic.Place.CategoryId)
		if err != nil {
			return nil, err
		}

		service := &model.ServiceResponse{
			ServiceId: beautyClinic.ID,
			Name:      beautyClinic.Name,
			PhotoLink: beautyClinic.PhotoLink,
			Rating:    beautyClinic.AvgRating,
			Address:   beautyClinic.Place.Address,
			Category:  category.Name,
			Price:     beautyClinic.Price,
			Hour:      beautyClinic.Place.Hour,
		}

		services = append(services, service)
	}

	return services, nil
}

func (s Service) GetById(id uint) (*model.ServiceDetailResponse, error) {
	service, err := s.serviceRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	var reviews []*entity.Review
	for _, review := range service.Reviews {
		if review.ServiceID == id {

			data := &entity.Review{
				ID:        review.ID,
				UserID:    review.UserID,
				User:      review.User,
				ServiceID: review.ServiceID,
				Rating:    review.Rating,
				Review:    review.Review,
				CreatedAt: review.CreatedAt,
				UpdatedAt: review.UpdatedAt,
			}
			reviews = append(reviews, data)
		}
	}

	response := &model.ServiceDetailResponse{
		Name:        service.Name,
		PhotoLink:   service.PhotoLink,
		Rating:      service.AvgRating,
		Address:     service.Place.Address,
		Description: service.Description,
		Price:       service.Price,
		Review:      reviews,
	}

	return response, nil
}

func (s Service) GetAllSpaMassage(filter model.FilterParam, page int) ([]*model.ServiceResponse, error) {
	limit := 5
	offset := (page - 1) * limit

	spaMassages, err := s.serviceRepo.GetAllFitnessCenter(filter, limit, offset)
	if err != nil {
		return nil, err
	}

	var services []*model.ServiceResponse
	for _, spaMassage := range spaMassages {
		category, err := s.categoryRepo.GetById(spaMassage.Place.CategoryId)
		if err != nil {
			return nil, err
		}

		service := &model.ServiceResponse{
			ServiceId: spaMassage.ID,
			Name:      spaMassage.Name,
			PhotoLink: spaMassage.PhotoLink,
			Rating:    spaMassage.AvgRating,
			Address:   spaMassage.Place.Address,
			Category:  category.Name,
			Price:     spaMassage.Price,
			Hour:      spaMassage.Place.Hour,
		}

		services = append(services, service)
	}

	return services, nil
}

func (s Service) GetAllSalon(filter model.FilterParam, page int) ([]*model.ServiceResponse, error) {
	limit := 5
	offset := (page - 1) * limit

	salons, err := s.serviceRepo.GetAllSalon(filter, limit, offset)
	if err != nil {
		return nil, err
	}

	var services []*model.ServiceResponse
	for _, salon := range salons {
		category, err := s.categoryRepo.GetById(salon.Place.CategoryId)
		if err != nil {
			return nil, err
		}

		service := &model.ServiceResponse{
			ServiceId: salon.ID,
			Name:      salon.Name,
			PhotoLink: salon.PhotoLink,
			Rating:    salon.AvgRating,
			Address:   salon.Place.Address,
			Category:  category.Name,
			Price:     salon.Price,
			Hour:      salon.Place.Hour,
		}

		services = append(services, service)
	}

	return services, nil
}

func (s Service) GetAllFitnessCenter(filter model.FilterParam, page int) ([]*model.ServiceResponse, error) {
	limit := 5
	offset := (page - 1) * limit

	fitnessCenters, err := s.serviceRepo.GetAllFitnessCenter(filter, limit, offset)
	if err != nil {
		return nil, err
	}

	var services []*model.ServiceResponse
	for _, fitnessCenter := range fitnessCenters {
		category, err := s.categoryRepo.GetById(fitnessCenter.Place.CategoryId)
		if err != nil {
			return nil, err
		}

		service := &model.ServiceResponse{
			ServiceId: fitnessCenter.ID,
			Name:      fitnessCenter.Name,
			PhotoLink: fitnessCenter.PhotoLink,
			Rating:    fitnessCenter.AvgRating,
			Address:   fitnessCenter.Place.Address,
			Category:  category.Name,
			Price:     fitnessCenter.Price,
			Hour:      fitnessCenter.Place.Hour,
		}

		services = append(services, service)
	}

	return services, nil
}

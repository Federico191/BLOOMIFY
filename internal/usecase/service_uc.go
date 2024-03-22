package usecase

import (
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
	"projectIntern/model"
)

type ServiceUCItf interface {
	GetById(id uint) (*model.ServiceDetailResponse, error)
	GetAllBeautyClinic(filter model.FilterParam, page int) ([]*model.ServiceResponse, error)
	GetAllSpaMassage(filter model.FilterParam, page int) ([]*model.ServiceResponse, error)
	GetAllSalon(filter model.FilterParam, page int) ([]*model.ServiceResponse, error)
	GetAllFitnessCenter(filter model.FilterParam, page int) ([]*model.ServiceResponse, error)
	GetByTopRate() ([]*model.ServiceResponseDashboard, error)
	GetByProblem(problemId uint) ([]*model.ServiceResponseDashboard, error)
}

type ServiceUC struct {
	serviceRepo  repository.ServiceRepoItf
	categoryRepo repository.CategoryRepoItf
	reviewRepo   repository.TreatmentReviewRepoItf
	userRepo     repository.UserRepoItf
}

func NewServiceUC(repo repository.ServiceRepoItf, categoryRepo repository.CategoryRepoItf) ServiceUCItf {
	return &ServiceUC{serviceRepo: repo, categoryRepo: categoryRepo}
}

func (s ServiceUC) GetAllBeautyClinic(filter model.FilterParam, page int) ([]*model.ServiceResponse, error) {
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

func (s ServiceUC) GetById(id uint) (*model.ServiceDetailResponse, error) {
	service, err := s.serviceRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	var reviews []*entity.TreatmentReview
	for _, review := range service.Reviews {
		if review.ServiceID == id {

			data := &entity.TreatmentReview{
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
		Name:         service.Name,
		PhotoLink:    service.PhotoLink,
		Rating:       service.AvgRating,
		Address:      service.Place.Address,
		Description:  service.Description,
		Price:        service.Price,
		ReviewerName: service.Reviews[0].User.FullName,
		Review:       service.Reviews[0].Review,
		ReviewRating: service.Reviews[0].Rating,
	}

	return response, nil
}

func (s ServiceUC) GetAllSpaMassage(filter model.FilterParam, page int) ([]*model.ServiceResponse, error) {
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

func (s ServiceUC) GetAllSalon(filter model.FilterParam, page int) ([]*model.ServiceResponse, error) {
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

func (s ServiceUC) GetAllFitnessCenter(filter model.FilterParam, page int) ([]*model.ServiceResponse, error) {
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

func (s ServiceUC) GetByTopRate() ([]*model.ServiceResponseDashboard, error) {
	services, err := s.serviceRepo.GetByTopRate()
	if err != nil {
		return nil, err
	}

	var serviceResponses []*model.ServiceResponseDashboard
	for _, service := range services {
		response := &model.ServiceResponseDashboard{
			ServiceId:   service.ID,
			Name:        service.Name,
			PhotoLink:   service.PhotoLink,
			PlaceName:   service.Place.Name,
			Description: service.Description,
		}

		serviceResponses = append(serviceResponses, response)
	}

	return serviceResponses, nil
}

func (s ServiceUC) GetByProblem(problemId uint) ([]*model.ServiceResponseDashboard, error) {
	services, err := s.serviceRepo.GetByProblem(problemId)
	if err != nil {
		return nil, err
	}

	var serviceResponses []*model.ServiceResponseDashboard
	for _, service := range services {
		response := &model.ServiceResponseDashboard{
			ServiceId:   service.ID,
			Name:        service.Name,
			PhotoLink:   service.PhotoLink,
			PlaceName:   service.Place.Name,
			Description: service.Description,
		}

		serviceResponses = append(serviceResponses, response)
	}

	return serviceResponses, nil
}

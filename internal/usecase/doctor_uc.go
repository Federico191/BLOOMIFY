package usecase

import (
	"github.com/google/uuid"
	"projectIntern/internal/repository"
	"projectIntern/model"
)

type DoctorUCItf interface {
	GetAll(filter model.FilterParam, page int) ([]*model.DoctorResponse, error)
	GetById(id uuid.UUID) (*model.DoctorDetailResponse, error)
}

type DoctorUC struct {
	doctorRepo repository.DoctorRepoItf
}

func NewDoctorUC(doctorRepo repository.DoctorRepoItf) DoctorUCItf {
	return &DoctorUC{doctorRepo: doctorRepo}
}

func (d DoctorUC) GetAll(filter model.FilterParam, page int) ([]*model.DoctorResponse, error) {
	limit := 5
	offset := (page - 1) * limit

	doctors, err := d.doctorRepo.GetAll(filter, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []*model.DoctorResponse
	for _, doctor := range doctors {
		response := &model.DoctorResponse{
			DocterID:     doctor.ID,
			DoctorName:   doctor.Name,
			Profession:   doctor.Profession.Name,
			Price:        doctor.Price,
			Rating:       doctor.Rating,
			Alumnus:      doctor.Alumnus,
			PracticeSite: doctor.PracticeSite,
			STRNumber:    doctor.STRNumber,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (d DoctorUC) GetById(id uuid.UUID) (*model.DoctorDetailResponse, error) {
	doctor, err := d.doctorRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	response := &model.DoctorDetailResponse{
		DocterId:     doctor.ID,
		DoctorName:   doctor.Name,
		Profession:   doctor.Profession.Name,
		Price:        doctor.Price,
		Rating:       doctor.Rating,
		ReviewerName: doctor.Reviews[0].User.FullName,
		Review:       doctor.Reviews[0].Review,
		ReviewRating: doctor.Reviews[0].Rating,
	}

	return response, nil
}

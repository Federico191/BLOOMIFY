package usecase

import (
	"github.com/google/uuid"
	"projectIntern/internal/repository"
	"projectIntern/model"
	"time"
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
			DoctorID:     doctor.ID,
			DoctorName:   doctor.Name,
			Profession:   doctor.Profession.Name,
			Age:          calculateAge(doctor.BirthDate),
			City:         doctor.City,
			PhotoLink:    doctor.PhotoLink,
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
		DoctorId:     doctor.ID,
		DoctorName:   doctor.Name,
		Profession:   doctor.Profession.Name,
		Price:        doctor.Price,
		Age:          calculateAge(doctor.BirthDate),
		City:         doctor.City,
		Rating:       doctor.Rating,
		PhotoLink:    doctor.PhotoLink,
		ReviewerName: doctor.Reviews[0].User.FullName,
		Review:       doctor.Reviews[0].Review,
		ReviewRating: doctor.Reviews[0].Rating,
	}

	return response, nil
}

func calculateAge(birthDate time.Time) int {
	now := time.Now()
	years := now.Year() - birthDate.Year()

	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		years--
	}

	return years
}

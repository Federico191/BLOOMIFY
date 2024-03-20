package model

import "github.com/google/uuid"

type DoctorResponse struct {
	DoctorID     uuid.UUID `json:"doctor_id"`
	DoctorName   string    `json:"doctor_name"`
	Profession   string    `json:"profession"`
	Age          int       `json:"age"`
	PhotoLink    string    `json:"photo_link"`
	Price        int       `json:"price"`
	Rating       float64   `json:"rating"`
	Alumnus      string    `json:"alumnus"`
	PracticeSite string    `json:"practice_site"`
	STRNumber    string    `json:"str_number"`
}

type DoctorDetailResponse struct {
	DoctorId     uuid.UUID `json:"doctor_id"`
	DoctorName   string    `json:"doctor_name"`
	Profession   string    `json:"profession"`
	Age          int       `json:"age"`
	Price        int       `json:"price"`
	Rating       float64   `json:"rating"`
	PhotoLink    string    `json:"photo_link"`
	ReviewerName string    `json:"reviewer_name"`
	Review       string    `json:"review"`
	ReviewRating int       `json:"review_rating"`
}

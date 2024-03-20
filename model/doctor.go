package model

import "github.com/google/uuid"

type DoctorResponse struct {
	DocterID     uuid.UUID `json:"docter_id"`
	DoctorName   string    `json:"doctor_name"`
	Profession   string    `json:"profession"`
	Price        int       `json:"price"`
	Rating       float64   `json:"rating"`
	Alumnus      string    `json:"alumnus"`
	PracticeSite string    `json:"practice_site"`
	STRNumber    string    `json:"str_number"`
}

type DoctorDetailResponse struct {
	DocterId     uuid.UUID `json:"docter_id"`
	DoctorName   string    `json:"doctor_name"`
	Profession   string    `json:"profession"`
	Price        int       `json:"price"`
	Rating       float64   `json:"rating"`
	ReviewerName string    `json:"reviewer_name"`
	Review       string    `json:"review"`
	ReviewRating int       `json:"review_rating"`
}

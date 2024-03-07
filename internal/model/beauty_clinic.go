package model

import (
	"github.com/google/uuid"
	"time"
)

type BeautyClinicRequest struct {
	Name      string  `json:"name" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	City      string  `json:"city" binding:"required"`
	Contact   string  `json:"contact" binding:"required"`
	Rating    float64 `json:"rating" binding:"required"`
	Hour      string  `json:"hour" binding:"required;"`
	PhotoLink string  `json:"photo_link"`
}

type BeautyClinicUpdate struct {
	ID        string  `json:"id" binding:"required;uuid"`
	Name      string  `json:"name" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	City      string  `json:"city" binding:"required"`
	Contact   string  `json:"contact" binding:"required"`
	Rating    float64 `json:"rating" binding:"required;numeric"`
	Hour      string  `json:"hour" binding:"required"`
	PhotoLink string  `json:"photo_link"`
}

type BeautyClinicResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	Contact   string    `json:"contact"`
	Rating    float64   `json:"rating"`
	Hour      string    `json:"hour"`
	PhotoLink string    `json:"photo_link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

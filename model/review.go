package model

import "github.com/google/uuid"

type ReviewRequest struct {
	PlaceId uint      `json:"place_id"`
	UserId  uuid.UUID `json:"user_id"`
	Review  string    `json:"name"`
	Rating  int       `json:"rating"`
}

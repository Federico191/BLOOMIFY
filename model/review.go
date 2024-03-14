package model

import "github.com/google/uuid"

type ReviewRequest struct {
	ServiceId uint      `json:"service_id"`
	UserId    uuid.UUID `json:"user_id"`
	Review    string    `json:"name"`
	Rating    int       `json:"rating"`
}

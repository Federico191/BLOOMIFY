package model

import "projectIntern/internal/entity"

type PlaceResponse struct {
	Place      entity.Place `json:"place"`
	Pagination Pagination   `json:"pagination"`
}

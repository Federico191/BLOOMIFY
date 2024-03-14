package model

type ServiceResponse struct {
	ServiceId uint    `json:"service_id"`
	Name      string  `json:"name"`
	PhotoLink string  `json:"photo_link"`
	Rating    float64 `json:"rating"`
	Address   string  `json:"address"`
	Category  string  `json:"category"`
	Price     int     `json:"price"`
	Hour      string  `json:"hour"`
}

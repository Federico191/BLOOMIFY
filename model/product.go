package model

type ProductResponse struct {
	Name        string  `json:"name"`
	Problem     string  `json:"problem"`
	Price       int     `json:"price"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	PhotoLink   string  `json:"photo_link"`
}

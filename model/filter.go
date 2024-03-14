package model

type FilterParam struct {
	City   string `form:"city"`
	Price  string `form:"price"`
	Rating string `form:"rating"`
}

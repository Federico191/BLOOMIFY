package model

type FilterParam struct {
	City     string `form:"city"`
	MinPrice int    `form:"min_price"`
	MaxPrice int    `form:"max_price"`
	Rating   int    `form:"rating"`
}

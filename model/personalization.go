package model

type PersonalizationReq struct {
	Answers []string `json:"answers" binding:"required"`
}

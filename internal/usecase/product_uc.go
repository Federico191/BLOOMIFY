package usecase

import (
	"log"
	"projectIntern/internal/repository"
	"projectIntern/model"
)

type ProductUCItf interface {
	GetByProblem(problemId uint) ([]*model.ProductResponse, error)
	GetByTopRate() ([]*model.ProductResponse, error)
}

type ProductUC struct {
	product repository.ProductRepoItf
}

func NewProductUC(product repository.ProductRepoItf) ProductUCItf {
	return &ProductUC{product: product}
}

func (p ProductUC) GetByProblem(problemId uint) ([]*model.ProductResponse, error) {
	products, err := p.product.GetByProblem(problemId)
	log.Println(products)
	if err != nil {
		return nil, err
	}

	var productResponses []*model.ProductResponse
	for _, product := range products {
		response := &model.ProductResponse{
			Name:        product.Name,
			Problem:     product.Problem.Name,
			Price:       product.Price,
			Description: product.Description,
			Rating:      product.Rating,
			PhotoLink:   product.PhotoLink,
		}

		productResponses = append(productResponses, response)
	}

	return productResponses, nil
}

func (p ProductUC) GetByTopRate() ([]*model.ProductResponse, error) {
	products, err := p.product.GetByTopRate()
	if err != nil {
		return nil, err
	}

	var productResponses []*model.ProductResponse
	for _, product := range products {
		response := &model.ProductResponse{
			Name:        product.Name,
			Problem:     product.Problem.Name,
			Price:       product.Price,
			Description: product.Description,
			Rating:      product.Rating,
			PhotoLink:   product.PhotoLink,
		}
		productResponses = append(productResponses, response)
	}

	return productResponses, nil
}

package usecase

import (
	"go-products-api/model"
	"go-products-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	return pu.repository.CreateProduct(product)
}

func (pu *ProductUsecase) GetProductById(id int) (*model.Product, error) {
	return pu.repository.GetProductById(id)
}

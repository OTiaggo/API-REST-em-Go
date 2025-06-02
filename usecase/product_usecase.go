package usecase

import (
	"go-api/model"
	"go-api/repository"
)

// Regra do GO: tem que começar com letra maiúscula se quiser que a estrutura seja visível em outros pacotes
type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

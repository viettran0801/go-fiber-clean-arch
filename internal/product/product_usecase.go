package product

import (
	"gofiber/domain"
)

type productUsecase struct {
	productRepository domain.ProductRepository
}

func NewProductUsecase(pr domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		productRepository: pr,
	}
}

func (p *productUsecase) Create(product *domain.Product) error {
	return p.productRepository.Create(product)
}

func (p *productUsecase) GetAll() (*[]domain.Product, error) {
	return p.productRepository.GetAll()
}

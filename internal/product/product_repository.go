package product

import (
	"gofiber/domain"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) Create(product *domain.Product) error {
	p.db.Create(product)
	return nil
}

func (p *productRepository) GetAll() (*[]domain.Product, error) {
	products := []domain.Product{}
	p.db.Find(&products)
	return &products, nil
}

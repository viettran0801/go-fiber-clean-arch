package domain

type Product struct {
	BaseModel
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

type ProductRepository interface {
	Create(product *Product) error
	GetAll() (*[]Product, error)
}

type ProductUsecase interface {
	Create(product *Product) error
	GetAll() (*[]Product, error)
}

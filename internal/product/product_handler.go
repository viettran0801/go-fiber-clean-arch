package product

import (
	"gofiber/domain"

	"github.com/gofiber/fiber/v2"
)

type ProductHandle struct {
	ProductUsecase domain.ProductUsecase
}

func NewProductHandle(r fiber.Router, pu domain.ProductUsecase) {
	handle := &ProductHandle{
		ProductUsecase: pu,
	}

	r.Get("/", handle.GetAllProduct)
	r.Post("/", handle.CreateProduct)
}

func (p *ProductHandle) GetAllProduct(c *fiber.Ctx) error {
	products, err := p.ProductUsecase.GetAll()
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"status": "success",
		"data":   products,
	})
}

func (p *ProductHandle) CreateProduct(c *fiber.Ctx) error {
	product := new(domain.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := p.ProductUsecase.Create(product); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"status":  "success",
		"product": product,
	})

}

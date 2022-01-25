package infrastructure

import (
	"gofiber/internal/product"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Run() {

	db, err := ConnectToDB()
	if err != nil {
		log.Fatal("Database connection failed: $s", err)
	}

	app := fiber.New(fiber.Config{
		AppName:      "Fiber clean arch",
		ServerHeader: "Fiber",
	})

	// Config global middlewares
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Too many request",
			})
		},
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	// Create repositories
	productRepository := product.NewProductRepository(db)

	// Create usecase
	productUsecase := product.NewProductUsecase(productRepository)

	// create endpoint API
	product.NewProductHandle(app.Group("/v1/products"), productUsecase)

	log.Fatal(app.Listen(":8000"))
}

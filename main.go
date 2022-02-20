package main

import (
	"code/config"
	"code/entity/product"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := config.GetDbConn()
	productRepo := product.NewProductRepo(db)
	productService := product.NewProductService(productRepo)
	productHandler := product.NewProductHandler(productService)

	app.Get("/", productHandler.Find)

	app.Listen(":3000")
}

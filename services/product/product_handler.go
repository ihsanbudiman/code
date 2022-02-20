package product

import "github.com/gofiber/fiber/v2"

type productHandler struct {
	productService ProductService
}

func NewProductHandler(productService ProductService) *productHandler {
	return &productHandler{productService: productService}
}

func (p *productHandler) Find(c *fiber.Ctx) error {
	product, err := p.productService.Find()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(product)
}

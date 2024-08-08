package controller

import (
	"github.com/123-zuleyha/e-commerce/models"
	"github.com/gofiber/fiber/v2"
)

// GetProducts tüm ürünleri listeleyen fonksiyon GET
func GetProducts(c *fiber.Ctx) error {
	products, err := models.GetAllProducts()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

// GetProduct belirli bir id'ye sahip ürünü listeleyen fonksiyon GET
// Param: id
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := models.GetProductByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Ürün bulunamadı.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

// CreateProduct yeni ürün oluşturmak için POST isteği oluşturur
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := models.CreateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

// UpdateProduct mevcut ürünü güncellemek için kullanılır. PUT isteklerini işler.
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := models.UpdateProduct(id, &product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

// DeleteProduct mevcut ürünü siler. DELETE isteklerini işler.
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := models.DeleteProduct(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ürün başarıyla silindi",
	})
}

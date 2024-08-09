package controller

import (
	"github.com/123-zuleyha/e-commerce/config"
	"github.com/123-zuleyha/e-commerce/models"
	"github.com/gofiber/fiber/v2"
)

// GetProducts tüm ürünleri listeleyen fonksiyon GET
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	config.DB.Find(&products)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"data":   products,
	})
}

// GetProduct belirli bir id'ye sahip ürünü listeleyen fonksiyon GET
// Param: id
func GetProductByID(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	var product models.Product
	config.DB.Find(&product, productID)

	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "product not found",
		})

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "succes",
			"data":   product,
		})

	}}



// CreateProduct yeni ürün oluşturmak için POST isteği oluşturur
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if product.Name =="" || product.Description == "" { //Eğer name ve description boşsa
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : "error",
			"message" : "name or description is empty" ,
		})
		config.DB.Create(&product) //veritananına ekleme yapıyoruz
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status" :"success" , 
			"data" : product ,
		})

	}

	

	
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

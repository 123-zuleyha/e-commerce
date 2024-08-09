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
			"data":   product ,
		})

	}



// CreateProduct yeni ürün oluşturmak için POST isteği oluşturur
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if product.Name ==""  { 
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : "error",
			"message" : "name  is empty" ,
		})
	if product.Description == ""  { 
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status" : "error" ,
				"message" : "description is empty" ,
			}) 
		}}

	if product.Price ==0  { 
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status" : "error",
					"message" : "price is empty or zero" ,
				})}
	if product.Stock ==0  { 
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"status" : "error",
						"message" : "stock is empty or zero" ,
					})}
	

     config.DB.Create(&product) //veritananına ekleme yapıyoruz
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status" :"success" , 
			"data" : product ,
		})

	}

	

	
}

// UpdateProduct mevcut ürünü güncellemek için kullanılır. PUT isteklerini işler.
func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	
	if reqProduct.Name !="" {
			product.Name = reqProduct.Name
		}

	if reqProduct.Description !="" {
		product.Description = reqProduct.Description
	}

	if reqProduct.Price > 0 && reqProduct.Price !=0 {
		product.Price = reqProduct.Price
	}

	if reqProduct.Stock > 0 && reqProduct.Stock !=0 {
		product.Stock = reqProduct.Stock
	}

    config.DB.Save(&product)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status" : "success" , 
		"data" : product ,
	})
	}
      
	

	
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

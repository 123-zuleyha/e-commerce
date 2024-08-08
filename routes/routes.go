package routes

import (
	"github.com/123-zuleyha/e-commerce/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	productRoutes := app.Group("/products")

	productRoutes.Get("/", controller.GetProducts)
	productRoutes.Get("/:id", controller.GetProduct) 
	productRoutes.Post("/", controller.CreateProduct)
	productRoutes.Put("/:id", controller.UpdateProduct) 
	productRoutes.Delete("/:id", controller.DeleteProduct) 
}
package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"toxicboy/adapters"
	"toxicboy/core/orders"
)

func SetupRoutes(app *fiber.App, db *gorm.DB){
	
	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := orders.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)
	app.Get("/order", orderHandler.FindAllOrders)
	app.Get("/order/:id", orderHandler.FindOrderById)
}
package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"toxicboy/adapters/http"
	"toxicboy/adapters/database"
	"toxicboy/core/orders"
)

func SetupOrderRoutes(app *fiber.App, db *gorm.DB){
	
	orderRepo := database.NewGormOrderRepository(db)
	orderService := orders.NewOrderService(orderRepo)
	orderHandler := http.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)
	app.Get("/order", orderHandler.FindAllOrders)
	app.Get("/order/:id", orderHandler.FindOrderById)
}
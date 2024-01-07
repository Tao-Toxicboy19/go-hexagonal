package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"toxicboy/core/orders"
	"toxicboy/routes"
	"toxicboy/config"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// Connection to PostgreSQL
	db, err := config.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}

	routes.SetupRoutes(app, db)
	
	db.AutoMigrate(&orders.Order{})

	app.Listen(":8080")
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"toxicboy/core/orders"
	"toxicboy/routes"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// Connection to PostgreSQL
	dsn := "user=postgres password=testpass123 dbname=go-hex port=3500 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	routes.SetupRoutes(app, db)
	
	db.AutoMigrate(&orders.Order{})

	app.Listen(":8080")
}

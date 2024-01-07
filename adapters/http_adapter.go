package adapters

import (
	"github.com/gofiber/fiber/v2"
	"toxicboy/core/orders"
)

type HttpOrderHandler struct {
	service orders.OrderService
}

func NewHttpOrderHandler(service orders.OrderService) *HttpOrderHandler {
	return &HttpOrderHandler{service: service}
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order orders.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/learning_grpc/services/common/genproto/orders"
	"github.com/yrnThiago/learning_grpc/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: orderService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *fiber.App) {
	router.Post("/orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(c *fiber.Ctx) error {
	var input orders.CreateOrderRequest
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "order body missing"})
	}

	output := &orders.Order{
		OrderID:    42,
		CustomerID: input.CustomerID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
	}

	err = h.ordersService.CreateOrder(context.Background(), output)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	return c.Status(fiber.StatusCreated).JSON(res)
}

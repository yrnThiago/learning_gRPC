package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	handler "github.com/yrnThiago/learning_grpc/services/orders/handler/orders"
	"github.com/yrnThiago/learning_grpc/services/orders/service"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}

func (s *HttpServer) Run() error {
	router := fiber.New()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Printf("Starting server on %s", s.addr)

	return router.Listen(s.addr)
}

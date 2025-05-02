package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/learning_grpc/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := fiber.New()

	conn := NewGRPCClient(":50051")
	defer conn.Close()

	router.Get("/orders", func(c *fiber.Ctx) error {
		client := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		_, err := client.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 24,
			ProductID:  3123,
			Quantity:   2,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		res, err := client.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerID: 42,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		return c.Status(200).JSON(res)
		return c.Render("orders", res.GetOrders(), ordersTemplate)
	})

	log.Println("Starting server on", s.addr)
	return router.Listen(s.addr)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`

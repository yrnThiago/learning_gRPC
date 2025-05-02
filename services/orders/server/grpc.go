package server

import (
	"log"
	"net"

	handler "github.com/yrnThiago/learning_grpc/services/orders/handler/orders"
	"github.com/yrnThiago/learning_grpc/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Printf("listening port on %s", ":50051")

	return grpcServer.Serve(listen)
}

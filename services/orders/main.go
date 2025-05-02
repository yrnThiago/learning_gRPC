package main

import "github.com/yrnThiago/learning_grpc/services/orders/server"

func main() {
	httpServer := server.NewHttpServer(":8000")
	go httpServer.Run()

	grpcServer := server.NewGRPCServer(":50051")
	grpcServer.Run()
}

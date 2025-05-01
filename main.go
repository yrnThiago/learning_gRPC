package main

import (
	"github.com/yrnThiago/learning_grpc/client"
	"github.com/yrnThiago/learning_grpc/server"
)

func main() {
	go server.Run()
	client.Run()
}

package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/yrnThiago/learning_grpc/pb"
)

func Run() {
	dial, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	defer dial.Close()

	userClient := pb.NewUserClient(dial)
	paymentClient := pb.NewPaymentClient(dial)

	user, err := userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "1",
		Age:  20,
		Name: "Test",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("new user created: %v\n", user)

	user, err = userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "2",
		Age:  22,
		Name: "Test2",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("new user created: %v\n", user)

	getUserResponse, err := userClient.GetUser(context.Background(), &pb.GetUserRequest{Id: "2"})
	if err != nil {
		return
	}

	fmt.Printf("User returned from GetUser method: %v\n", getUserResponse)

	fmt.Printf("Generating pix code for order id %s\n", "123")

	generatedPayment, err := paymentClient.GeneratePix(context.Background(), &pb.GeneratePixRequest{OrderId: "123"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("pix code: %s (order: %s)\n", generatedPayment.PixCode, generatedPayment.OrderId)
}

package main

import (
	"context"
	"fmt"
	pb "go-task/auth-service/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return
	}

	client := pb.NewAuthServiceClient(dial)
	user, err := client.LoginUser(context.Background(), &pb.LoginRequest{
		Email:    "johndoe@example.com",
		Password: "password",
	})
	if err != nil {
		return
	}

	fmt.Println(user)
}

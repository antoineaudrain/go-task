package main

import (
	"context"
	"fmt"
	pb "go-task/user/api"
	"google.golang.org/grpc"
)

func main() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return
	}

	client := pb.NewUserServiceClient(dial)
	//user, err := client.Create(context.Background(), &pb.CreateRequest{
	//	Email:    "johndoe@example.com",
	//	Password: "password",
	//	FullName: "John Doe",
	//})

	user, err := client.Login(context.Background(), &pb.LoginRequest{
		Email:    "johndoe@example.com",
		Password: "password",
	})

	//refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYxNTQzNjEsInVzZXJJZCI6Ijk0NTkyYmE3LTFmNTctNGMwYy1iNGQ0LWZjNDg2ODhhZmM3OCJ9.u49dcgXtHRS0edCbbrrq9sS7CEbglnr7gvOXvp3N2TM"
	//user, err := client.RefreshToken(context.Background(), &pb.RefreshTokenRequest{
	//	RefreshToken: refreshToken,
	//})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
}

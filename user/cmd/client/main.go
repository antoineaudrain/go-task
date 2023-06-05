package main

import (
	"context"
	"fmt"
	pb "go-task/user/pkg/pb"
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

	//refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYwNzIxNDEsInVzZXJfaWQiOiJjNmM5ZThhNS0wNWYzLTQ1OWQtOTBlNS1iYWRmYWJkNjRkYzUifQ.5y01xPxawpPpG92t9rN4ljszWjlUL6vRAaepZFPN594"
	//user, err := client.RefreshToken(context.Background(), &pb.RefreshTokenRequest{
	//	RefreshToken: refreshToken,
	//})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
}

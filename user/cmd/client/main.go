package main

import (
	"context"
	"fmt"
	pb "go-task/user/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return
	}

	client := pb.NewUserServiceClient(dial)

	//user, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{
	//	Email:    "johndoe@example.com",
	//	Password: "password",
	//	FullName: "John Doe",
	//})

	user, err := client.Login(context.Background(), &pb.LoginRequest{
		Email:    "johndoe@example.com",
		Password: "password",
	})

	//refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYyMzg1MTcsInN1YiI6IjQ5NmI5MWQyLWFlOTMtNDczZi04OTIzLTYyZDM5NTFkYmNjMCIsInRva2VuVHlwZSI6InJlZnJlc2gifQ.dF5hqFiCu7vQm8JYCU2gVcwMwC9unq9aW5nrNDFMCTY"
	//user, err := client.RefreshAccessToken(context.Background(), &pb.RefreshAccessTokenRequest{
	//	RefreshToken: refreshToken,
	//})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
}

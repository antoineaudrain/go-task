package main

import (
	"context"
	"fmt"
	pb "go-task/workspace/api"
	"google.golang.org/grpc"
)

func main() {
	dial, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		return
	}

	client := pb.NewWorkspaceServiceClient(dial)
	user, err := client.Create(context.Background(), &pb.CreateRequest{
		Name: "My Workspace",
	})
	if err != nil {
		return
	}

	fmt.Println(user)
}

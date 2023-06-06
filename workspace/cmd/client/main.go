package main

import (
	"context"
	"fmt"
	pb "go-task/workspace/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Connection failed: %v", err)
		}
	}(conn)

	client := pb.NewWorkspaceServiceClient(conn)

	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYwNTYxMzcsInVzZXJJZCI6ImQ2M2M1NGE5LTRlOTctNDQzMC1iMDg0LWQ2ZmI4YTM3YjNhNyJ9.DSstGMaMPg1yW3T_715gKcIOk1H7gItYK6yy4jihryA"
	res, err := client.Create(
		metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+accessToken),
		&pb.CreateRequest{
			Name: "My Workspace",
		})
	if err != nil {
		log.Fatalf("Create request failed: %v", err)
	}

	fmt.Println(res)
}

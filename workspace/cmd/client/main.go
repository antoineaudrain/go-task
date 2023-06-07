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

	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYxNTMwNTgsInN1YiI6IjQ5NmI5MWQyLWFlOTMtNDczZi04OTIzLTYyZDM5NTFkYmNjMCIsInRva2VuVHlwZSI6ImFjY2VzcyJ9.umNIFdtb0IleD3z2EBBjP7x7bumE6mEgqAm3u9R1qGw"
	res, err := client.CreateWorkspace(
		metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+accessToken),
		&pb.CreateWorkspaceRequest{
			Name: "My Workspace",
		})
	if err != nil {
		log.Fatalf("Create request failed: %v", err)
	}

	fmt.Println(res)
}

package main

import (
	"context"
	"fmt"
	pb "github.com/antoineaudrain/go-task/invitation/api"
	"github.com/antoineaudrain/go-task/invitation/pkg/client"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	workspaceServiceClient, err := client.NewClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create invitation client: %v", err)
	}

	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgyOTE3MTIsInN1YiI6Ijk5MTU1NjUwLTc3YzAtNGY5ZS1iMzJjLTAxNTM3NmVjMTczZiIsInRva2VuVHlwZSI6ImFjY2VzcyJ9.vd7gUVETpUOeYZEnsWNRSjJzK5TJeCzJ4wx2yaFSkEQ"

	res, err := workspaceServiceClient.SendWorkspaceInvitation(
		metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+accessToken),
		&pb.SendWorkspaceInvitationRequest{
			Email: "johndoe@example.com",
		})

	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	fmt.Println(res)
}

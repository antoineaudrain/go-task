package main

import (
	"context"
	"fmt"
	pb "github.com/antoineaudrain/go-task/workspace/api"
	"github.com/antoineaudrain/go-task/workspace/pkg/client"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	workspaceServiceClient, err := client.NewClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create workspace client: %v", err)
	}

	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgyOTE3MTIsInN1YiI6Ijk5MTU1NjUwLTc3YzAtNGY5ZS1iMzJjLTAxNTM3NmVjMTczZiIsInRva2VuVHlwZSI6ImFjY2VzcyJ9.vd7gUVETpUOeYZEnsWNRSjJzK5TJeCzJ4wx2yaFSkEQ"

	//res, err := workspaceServiceClient.CreateWorkspace(
	//	metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+accessToken),
	//	&pb.CreateWorkspaceRequest{
	//		Name: "My Workspace",
	//	})

	//res, err := workspaceServiceClient.UpdateWorkspace(
	//	metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+accessToken),
	//	&pb.UpdateWorkspaceRequest{
	//		WorkspaceId: "1fcde286-3a3a-480d-b509-9a4ecc366ba3",
	//		Name:        "My Other Workspace",
	//	})

	//res, err := workspaceServiceClient.GetWorkspace(
	//	metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+accessToken),
	//	&pb.GetWorkspaceRequest{
	//		WorkspaceId: "267526bb-6d5a-430c-9932-1a9055cc0607",
	//	})

	//res, err := workspaceServiceClient.DeleteWorkspace(
	//	metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+accessToken),
	//	&pb.DeleteWorkspaceRequest{
	//		WorkspaceId: "1fcde286-3a3a-480d-b509-9a4ecc366ba3",
	//	})

	res, err := workspaceServiceClient.ListWorkspaceMembers(
		metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+accessToken),
		&pb.ListWorkspaceMembersRequest{
			WorkspaceId: "267526bb-6d5a-430c-9932-1a9055cc0607",
		})

	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	fmt.Println(res)
}

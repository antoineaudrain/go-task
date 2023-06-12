package main

import (
	"context"
	"fmt"
	userPb "go-task/user/pkg/proto"
	pb "go-task/workspace/pkg/proto"
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

	conn2, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func(conn2 *grpc.ClientConn) {
		err := conn2.Close()
		if err != nil {
			log.Fatalf("Connection failed: %v", err)
		}
	}(conn2)

	userClient := userPb.NewUserServiceClient(conn2)

	user, err := userClient.Login(context.Background(), &userPb.LoginRequest{
		Email:    "johndoe@example.com",
		Password: "password",
	})

	client := pb.NewWorkspaceServiceClient(conn)

	//res, err := client.CreateWorkspace(
	//	metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+user.GetAccessToken()),
	//	&pb.CreateWorkspaceRequest{
	//		Name: "My Other Workspace",
	//	},
	//)

	res, err := client.ListWorkspaces(
		metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+user.GetAccessToken()),
		&pb.ListWorkspaceRequest{},
	)

	//res, err := client.GetWorkspace(
	//	metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+user.GetAccessToken()),
	//	&pb.GetWorkspaceRequest{
	//		WorkspaceID: "1cde76e2-4baa-4d6e-8ab2-a1d314e001b3",
	//	},
	//)

	//res, err := client.UpdateWorkspace(
	//	metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+user.GetAccessToken()),
	//	&pb.UpdateWorkspaceRequest{
	//		WorkspaceID: "1cde76e2-4baa-4d6e-8ab2-a1d314e001b3",
	//		Name:        "My Workspace (to be deleted)",
	//	},
	//)

	//res, err := client.DeleteWorkspace(
	//	metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+user.GetAccessToken()),
	//	&pb.DeleteWorkspaceRequest{
	//		WorkspaceID: "1cde76e2-4baa-4d6e-8ab2-a1d314e001b3",
	//	},
	//)

	if err != nil {
		log.Fatalf("Create request failed: %v", err)
	}

	fmt.Println(res)
}

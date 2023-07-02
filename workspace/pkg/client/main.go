package client

import (
	pb "github.com/antoineaudrain/go-task/workspace/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(addr string) (pb.WorkspaceServiceClient, error) {
	creds := insecure.NewCredentials()

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	client := pb.NewWorkspaceServiceClient(conn)
	return client, nil
}

package client

import (
	pb "github.com/antoineaudrain/go-task/invitation/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(addr string) (pb.InvitationServiceClient, error) {
	creds := insecure.NewCredentials()

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	client := pb.NewInvitationServiceClient(conn)
	return client, nil
}
